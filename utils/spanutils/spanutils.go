package spanutils

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/go-orion/Orion/utils"
	"github.com/go-orion/Orion/utils/log"
	newrelic "github.com/newrelic/go-agent"
	opentracing "github.com/opentracing/opentracing-go"
	otext "github.com/opentracing/opentracing-go/ext"
	"go.elastic.co/apm"
	"google.golang.org/grpc/metadata"
)

// TracingSpan defines an interface for implementing a tracing span
type TracingSpan interface {
	End()
	Finish()
	SetTag(key string, value interface{})
	SetQuery(query string)
	SetError(msg string)
}

type tracingSpan struct {
	openSpan        opentracing.Span
	datastore       bool
	external        bool
	dataSegment     newrelic.DatastoreSegment
	externalSegment newrelic.ExternalSegment
	segment         newrelic.Segment
	elasticSpan     *apm.Span
}

func (span *tracingSpan) End() {
	if span == nil {
		// dont panic when called against a nil span
		return
	}
	span.openSpan.Finish()

	if span.datastore {
		span.dataSegment.End()
	} else if span.external {
		span.externalSegment.End()
	} else {
		span.segment.End()
	}

	if span.elasticSpan != nil {
		span.elasticSpan.End()
	}
}

func (span *tracingSpan) Finish() {
	span.End()
}

func (span *tracingSpan) SetTag(key string, value interface{}) {
	if span == nil {
		// dont panic when called against a nil span
		return
	}
	span.openSpan.SetTag(key, value)
}

func (span *tracingSpan) SetQuery(query string) {
	if span == nil {
		// dont panic when called against a nil span
		return
	}
	span.openSpan.SetTag("query", query)
	if span.datastore {
		span.dataSegment.ParameterizedQuery = query
	}
}

func (span *tracingSpan) SetError(msg string) {
	if span == nil {
		// dont panic when called against a nil span
		return
	}
	if msg != "" {
		span.openSpan.SetTag("error", msg)
	}
}

//NewInternalSpan starts a span for tracing internal actions
func NewInternalSpan(ctx context.Context, name string) (TracingSpan, context.Context) {
	zip, ctx := opentracing.StartSpanFromContext(ctx, name)
	txn := utils.GetNewRelicTransactionFromContext(ctx)
	seg := newrelic.Segment{
		StartTime: newrelic.StartSegmentNow(txn),
		Name:      name,
	}
	eSpan, ctx := apm.StartSpan(ctx, name, "internal")
	return &tracingSpan{
		openSpan:    zip,
		segment:     seg,
		elasticSpan: eSpan,
	}, ctx
}

//NewDatastoreSpan starts a span for tracing data store actions
func NewDatastoreSpan(ctx context.Context, name string, datastore string) (TracingSpan, context.Context) {
	if !strings.HasPrefix(name, datastore) {
		name = datastore + name
	}
	zip, ctx := opentracing.StartSpanFromContext(ctx, name)
	zip.SetTag("store", datastore)
	txn := utils.GetNewRelicTransactionFromContext(ctx)
	seg := newrelic.DatastoreSegment{
		StartTime: newrelic.StartSegmentNow(txn),
		Product:   newrelic.DatastoreProduct(datastore),
		Operation: name,
	}
	eSpan, ctx := apm.StartSpan(ctx, name, "datastore")
	return &tracingSpan{
		openSpan:    zip,
		dataSegment: seg,
		datastore:   true,
		elasticSpan: eSpan,
	}, ctx
}

func buildExternalSpan(ctx context.Context, name string, url string) (*tracingSpan, context.Context) {
	ctx, zip := ClientSpan(name, ctx)
	zip.SetTag("url", url)
	txn := utils.GetNewRelicTransactionFromContext(ctx)
	seg := newrelic.ExternalSegment{
		StartTime: newrelic.StartSegmentNow(txn),
		URL:       url,
	}
	eSpan, ctx := apm.StartSpan(ctx, name, "external")
	return &tracingSpan{
		openSpan:        zip,
		externalSegment: seg,
		external:        true,
		elasticSpan:     eSpan,
	}, ctx
}

//NewExternalSpan starts a span for tracing external actions
func NewExternalSpan(ctx context.Context, name string, url string) (TracingSpan, context.Context) {
	return buildExternalSpan(ctx, name, url)
}

//NewHTTPExternalSpan starts a span for tracing external HTTP actions
func NewHTTPExternalSpan(ctx context.Context, name string, url string, hdr http.Header) (TracingSpan, context.Context) {
	s, ctx := buildExternalSpan(ctx, name, url)
	traceHTTPHeaders(ctx, s.openSpan, hdr)
	return s, ctx
}

func traceHTTPHeaders(ctx context.Context, sp opentracing.Span, hdr http.Header) {
	// Transmit the span's TraceContext as HTTP headers on our
	// outbound request.
	opentracing.GlobalTracer().Inject(
		sp.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(hdr))
}

// A type that conforms to opentracing.TextMapReader and
// opentracing.TextMapWriter.
type metadataReaderWriter struct {
	*metadata.MD
}

func (w metadataReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	if strings.HasSuffix(key, "-bin") {
		val = string(base64.StdEncoding.EncodeToString([]byte(val)))
	}
	(*w.MD)[key] = append((*w.MD)[key], val)
}

func (w metadataReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vals := range *w.MD {
		for _, v := range vals {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

//ClientSpan starts a new client span linked to the existing spans if any are found
func ClientSpan(operationName string, ctx context.Context) (context.Context, opentracing.Span) {
	tracer := opentracing.GlobalTracer()
	var clientSpan opentracing.Span
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		clientSpan = tracer.StartSpan(
			operationName,
			opentracing.ChildOf(parentSpan.Context()),
		)
	} else {
		clientSpan = tracer.StartSpan(operationName)
	}
	otext.SpanKindRPCClient.Set(clientSpan)
	ctx = opentracing.ContextWithSpan(ctx, clientSpan)
	return ctx, clientSpan
}

func GRPCTracingSpan(operationName string, ctx context.Context) context.Context {
	tracer := opentracing.GlobalTracer()
	// Retrieve gRPC metadata.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	if span := opentracing.SpanFromContext(ctx); span != nil {
		// There's nothing we can do with an error here.
		if err := tracer.Inject(span.Context(), opentracing.TextMap, metadataReaderWriter{&md}); err != nil {
			log.Info(ctx, "err", err, "component", "spanutils")
		}
	}

	var span opentracing.Span
	wireContext, err := tracer.Extract(opentracing.TextMap, metadataReaderWriter{&md})
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		log.Info(ctx, "err", err, "component", "spanutils")
	}
	span = tracer.StartSpan(operationName, otext.RPCServerOption(wireContext))
	ctx = opentracing.ContextWithSpan(ctx, span)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}
