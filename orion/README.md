# orion
`import "github.com/go-orion/Orion/orion"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package orion is a small lightweight framework written around grpc with the aim to shorten time to build microservices

Source code for Orion can be found at <a href="https://github.com/go-orion/Orion">https://github.com/go-orion/Orion</a>

It is derived from 'Framework' a small microservices framework written and used inside <a href="https://carousell.com">https://carousell.com</a>, It comes with a number of sensible defaults such as zipkin tracing, hystrix, live reload of configuration, etc.

### Why Orion
Orion uses protocol-buffers definitions (<a href="https://developers.google.com/protocol-buffers/docs/proto3">https://developers.google.com/protocol-buffers/docs/proto3</a>) using gRPC and orion proto plugin as base for building services.

Using proto definitions as our service base allows us to define clean contracts that everyone can understand and enables auto generation of client code.

You define your services as a proto definition, for example

	service SimpleService{
		rpc Echo (EchoRequest) returns (EchoResponse){
		}
	}
	message EchoRequest {
		string msg = 1;
	}
	message EchoResponse {
		string msg = 1;
	}

The above definition represents a service named 'SimpleService' which accepts 'EchoRequest' and returns 'EchoResponse' at 'Echo' endpoint.

After you have generated the code from protoc (using grpc and orion plugin), you need to implement the server interface generated by gRPC, orion uses this service definition and enables HTTP/gRPC calls to be made to the same service implementation.

### How do i use it
Lets go through the example defined at <a href="https://github.com/go-orion/Orion/tree/master/example/simple">https://github.com/go-orion/Orion/tree/master/example/simple</a> , It covers the minimum required implementation for orion. It has the following structure

	.
	├── cmd
	│   ├── client
	│   │   └── client.go
	│   └── server
	│       └── server.go
	├── service
	│   └── service.go
	└── simple_proto
		├── generate.sh
		└── simple.proto

First we define the proto 'simple_proto/simple.proto' as

	syntax = "proto3";
	
	package simple_proto;
	
	service SimpleService{
		rpc Echo (EchoRequest) returns (EchoResponse){
		}
	}
	
	message EchoRequest {
		string msg = 1;
	}
	
	message EchoResponse {
		string msg = 1;
	}

The above definition represents a service named 'SimpleService' which accepts 'EchoRequest' and returns 'EchoResponse' at 'Echo' endpoint.

Now we can execute 'generate.sh' which contains

	protoc -I . simple.proto --go_out=plugins=grpc:. --orion_out=.

Running this command generates the following file in the simple_proto directory:

	.
	├── generate.sh
	├── simple.pb.go
	├── simple.proto
	└── simple.proto.orion.pb.go

This contains:

	All the protocol buffer code to populate, serialize, and retrieve our request and response message types (simple.pb.go)
	An interface type (or stub) for clients to call with the methods defined in the SimpleService (simple.pb.go)
	An interface type for servers to implement, also with the methods defined in the SimpleService (simple.pb.go)
	Registration function for Orion (simple.proto.orion.pb.go)

### Whats Incuded
Orion comes included with.

	Hystrix (<a href="http://github.com/afex/hystrix-go">http://github.com/afex/hystrix-go</a>)
	Zipkin (<a href="http://github.com/opentracing/opentracing-go">http://github.com/opentracing/opentracing-go</a>)
	NewRelic (<a href="http://github.com/newrelic/go-agent">http://github.com/newrelic/go-agent</a>)
	Prometheus (<a href="http://github.com/grpc-ecosystem/go-grpc-prometheus">http://github.com/grpc-ecosystem/go-grpc-prometheus</a>)
	Pprof (<a href="https://golang.org/pkg/net/http/pprof/">https://golang.org/pkg/net/http/pprof/</a>))
	Configuration (<a href="http://github.com/spf13/viper">http://github.com/spf13/viper</a>)
	Live Configuration Reload (<a href="http://github.com/go-orion/Orion/utils/listenerutils">http://github.com/go-orion/Orion/utils/listenerutils</a>)
	And much more...

### Getting Started
First follow the install guide at <a href="https://github.com/go-orion/Orion/blob/master/README.md">https://github.com/go-orion/Orion/blob/master/README.md</a>

## <a name="pkg-imports">Imported Packages</a>

- [github.com/afex/hystrix-go/hystrix](https://godoc.org/github.com/afex/hystrix-go/hystrix)
- [github.com/afex/hystrix-go/hystrix/metric_collector](https://godoc.org/github.com/afex/hystrix-go/hystrix/metric_collector)
- [github.com/afex/hystrix-go/plugins](https://godoc.org/github.com/afex/hystrix-go/plugins)
- [github.com/go-kit/kit/log](https://godoc.org/github.com/go-kit/kit/log)
- [github.com/go-orion/Orion/orion/handlers](./handlers)
- [github.com/go-orion/Orion/orion/handlers/grpc](./handlers/grpc)
- [github.com/go-orion/Orion/orion/handlers/http](./handlers/http)
- [github.com/go-orion/Orion/utils](./../utils)
- [github.com/go-orion/Orion/utils/errors/notifier](./../utils/errors/notifier)
- [github.com/go-orion/Orion/utils/listenerutils](./../utils/listenerutils)
- [github.com/go-orion/Orion/utils/log](./../utils/log)
- [github.com/grpc-ecosystem/go-grpc-prometheus](https://godoc.org/github.com/grpc-ecosystem/go-grpc-prometheus)
- [github.com/newrelic/go-agent](https://godoc.org/github.com/newrelic/go-agent)
- [github.com/opentracing/opentracing-go](https://godoc.org/github.com/opentracing/opentracing-go)
- [github.com/openzipkin/zipkin-go-opentracing](https://godoc.org/github.com/openzipkin/zipkin-go-opentracing)
- [github.com/prometheus/client_golang/prometheus/promhttp](https://godoc.org/github.com/prometheus/client_golang/prometheus/promhttp)
- [github.com/spf13/viper](https://godoc.org/github.com/spf13/viper)
- [google.golang.org/grpc](https://godoc.org/google.golang.org/grpc)

## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func AddConfigPath(path ...string)](#AddConfigPath)
* [func RegisterDecoder(svr Server, serviceName, method string, decoder Decoder)](#RegisterDecoder)
* [func RegisterDefaultDecoder(svr Server, serviceName string, decoder Decoder)](#RegisterDefaultDecoder)
* [func RegisterDefaultEncoder(svr Server, serviceName string, encoder Encoder)](#RegisterDefaultEncoder)
* [func RegisterEncoder(svr Server, serviceName, method, httpMethod, path string, encoder Encoder)](#RegisterEncoder)
* [func RegisterEncoders(svr Server, serviceName, method string, httpMethod []string, path string, encoder Encoder)](#RegisterEncoders)
* [func RegisterHandler(svr Server, serviceName, method string, path string, handler HTTPHandler)](#RegisterHandler)
* [func RegisterMethodOption(svr Server, serviceName, method, option string)](#RegisterMethodOption)
* [func RegisterMiddleware(svr Server, serviceName, method string, middleware ...string)](#RegisterMiddleware)
* [func ResetConfigPath()](#ResetConfigPath)
* [type Config](#Config)
  * [func BuildDefaultConfig(name string) Config](#BuildDefaultConfig)
* [type Decoder](#Decoder)
* [type DefaultServerImpl](#DefaultServerImpl)
  * [func (d \*DefaultServerImpl) AddDecoder(serviceName, method string, decoder handlers.Decoder)](#DefaultServerImpl.AddDecoder)
  * [func (d \*DefaultServerImpl) AddDefaultDecoder(serviceName string, decoder Decoder)](#DefaultServerImpl.AddDefaultDecoder)
  * [func (d \*DefaultServerImpl) AddDefaultEncoder(serviceName string, encoder Encoder)](#DefaultServerImpl.AddDefaultEncoder)
  * [func (d \*DefaultServerImpl) AddEncoder(serviceName, method string, httpMethod []string, path string, encoder handlers.Encoder)](#DefaultServerImpl.AddEncoder)
  * [func (d \*DefaultServerImpl) AddHTTPHandler(serviceName string, method string, path string, handler handlers.HTTPHandler)](#DefaultServerImpl.AddHTTPHandler)
  * [func (d \*DefaultServerImpl) AddInitializers(ins ...Initializer)](#DefaultServerImpl.AddInitializers)
  * [func (d \*DefaultServerImpl) AddMiddleware(serviceName string, method string, middlewares ...string)](#DefaultServerImpl.AddMiddleware)
  * [func (d \*DefaultServerImpl) AddOption(serviceName, method, option string)](#DefaultServerImpl.AddOption)
  * [func (d \*DefaultServerImpl) GetConfig() map[string]interface{}](#DefaultServerImpl.GetConfig)
  * [func (d \*DefaultServerImpl) GetOrionConfig() Config](#DefaultServerImpl.GetOrionConfig)
  * [func (d \*DefaultServerImpl) RegisterService(sd \*grpc.ServiceDesc, sf interface{}) error](#DefaultServerImpl.RegisterService)
  * [func (d \*DefaultServerImpl) Start()](#DefaultServerImpl.Start)
  * [func (d \*DefaultServerImpl) Stop(timeout time.Duration) error](#DefaultServerImpl.Stop)
  * [func (d \*DefaultServerImpl) Wait() error](#DefaultServerImpl.Wait)
* [type Encoder](#Encoder)
* [type FactoryParams](#FactoryParams)
* [type HTTPHandler](#HTTPHandler)
* [type HystrixConfig](#HystrixConfig)
  * [func BuildDefaultHystrixConfig() HystrixConfig](#BuildDefaultHystrixConfig)
* [type Initializer](#Initializer)
  * [func ErrorLoggingInitializer() Initializer](#ErrorLoggingInitializer)
  * [func HystrixInitializer() Initializer](#HystrixInitializer)
  * [func NewRelicInitializer() Initializer](#NewRelicInitializer)
  * [func PprofInitializer() Initializer](#PprofInitializer)
  * [func PrometheusInitializer() Initializer](#PrometheusInitializer)
  * [func ZipkinInitializer() Initializer](#ZipkinInitializer)
* [type NewRelicConfig](#NewRelicConfig)
  * [func BuildDefaultNewRelicConfig() NewRelicConfig](#BuildDefaultNewRelicConfig)
* [type Server](#Server)
  * [func GetDefaultServer(name string) Server](#GetDefaultServer)
  * [func GetDefaultServerWithConfig(config Config) Server](#GetDefaultServerWithConfig)
* [type ServiceFactory](#ServiceFactory)
* [type ServiceFactoryV2](#ServiceFactoryV2)
  * [func ToServiceFactoryV2(sf interface{}) (ServiceFactoryV2, error)](#ToServiceFactoryV2)
* [type ZipkinConfig](#ZipkinConfig)
  * [func BuildDefaultZipkinConfig() ZipkinConfig](#BuildDefaultZipkinConfig)

#### <a name="pkg-files">Package files</a>
[config.go](./config.go) [core.go](./core.go) [documetation.go](./documetation.go) [external.go](./external.go) [initializer.go](./initializer.go) [types.go](./types.go) [utils.go](./utils.go) 

## <a name="pkg-constants">Constants</a>
``` go
const (
    //ProtoGenVersion1_0 is the version of protoc-gen-orion plugin compatible with current code base
    ProtoGenVersion1_0 = true
    //BANNER is the orion banner text
    BANNER = `
  ___  ____  ___ ___  _   _
 / _ \|  _ \|_ _/ _ \| \ | |
| | | | |_) || | | | |  \| |
| |_| |  _ < | | |_| | |\  |
 \___/|_| \_\___\___/|_| \_|
                            `
)
```

## <a name="pkg-variables">Variables</a>
``` go
var (
    //ErrNil when the passed argument is nil
    ErrNil = errors.New("nil argument passed")
    //ErrNotServiceFactory when passed argument is not a service factory
    ErrNotServiceFactory = errors.New("you need to pass either a ServiceFactory or ServiceFactoryV2")
)
```
``` go
var (
    //DefaultInitializers are the initializers applied by orion as default
    DefaultInitializers = []Initializer{
        HystrixInitializer(),
        ZipkinInitializer(),
        NewRelicInitializer(),
        PrometheusInitializer(),
        PprofInitializer(),
        ErrorLoggingInitializer(),
    }
)
```

## <a name="AddConfigPath">func</a> [AddConfigPath](./config.go#L170)
``` go
func AddConfigPath(path ...string)
```
AddConfigPath adds a config path from where orion tries to read config values

## <a name="RegisterDecoder">func</a> [RegisterDecoder](./external.go#L41)
``` go
func RegisterDecoder(svr Server, serviceName, method string, decoder Decoder)
```
RegisterDecoder allows for registering an HTTP request decoder to a method
Note: this is normally called from protoc-gen-orion autogenerated files

## <a name="RegisterDefaultDecoder">func</a> [RegisterDefaultDecoder](./external.go#L33)
``` go
func RegisterDefaultDecoder(svr Server, serviceName string, decoder Decoder)
```
RegisterDefaultDecoder allows for registering an HTTP request decoder to arbitrary urls for the entire service
Note: this is normally called from protoc-gen-orion autogenerated files

## <a name="RegisterDefaultEncoder">func</a> [RegisterDefaultEncoder](./external.go#L25)
``` go
func RegisterDefaultEncoder(svr Server, serviceName string, encoder Encoder)
```
RegisterDefaultEncoder allows for registering an HTTP request encoder to arbitrary urls for the entire service
Note: this is normally called from protoc-gen-orion autogenerated files

## <a name="RegisterEncoder">func</a> [RegisterEncoder](./external.go#L9)
``` go
func RegisterEncoder(svr Server, serviceName, method, httpMethod, path string, encoder Encoder)
```
RegisterEncoder allows for registering an HTTP request encoder to arbitrary urls
Note: this is normally called from protoc-gen-orion autogenerated files

## <a name="RegisterEncoders">func</a> [RegisterEncoders](./external.go#L17)
``` go
func RegisterEncoders(svr Server, serviceName, method string, httpMethod []string, path string, encoder Encoder)
```
RegisterEncoders allows for registering an HTTP request encoder to arbitrary urls
Note: this is normally called from protoc-gen-orion autogenerated files

## <a name="RegisterHandler">func</a> [RegisterHandler](./external.go#L49)
``` go
func RegisterHandler(svr Server, serviceName, method string, path string, handler HTTPHandler)
```
RegisterHandler allows registering an HTTP handler for a given path
Note: this is normally called from protoc-gen-orion autogenerated files

## <a name="RegisterMethodOption">func</a> [RegisterMethodOption](./external.go#L57)
``` go
func RegisterMethodOption(svr Server, serviceName, method, option string)
```
RegisterMethodOption allows for registering an handler option to a particular method
Note: this is normally called from protoc-gen-orion autogenerated files

## <a name="RegisterMiddleware">func</a> [RegisterMiddleware](./external.go#L65)
``` go
func RegisterMiddleware(svr Server, serviceName, method string, middleware ...string)
```
RegisterMiddleware allows for registering  middlewares to a particular method
Note: this is normally called from protoc-gen-orion autogenerated files

## <a name="ResetConfigPath">func</a> [ResetConfigPath](./config.go#L178)
``` go
func ResetConfigPath()
```
ResetConfigPath resets the configuration paths

## <a name="Config">type</a> [Config](./config.go#L18-L52)
``` go
type Config struct {
    //OrionServerName is the name of this orion server that is tracked
    OrionServerName string
    // GRPCOnly tells orion not to build HTTP/1.1 server and only initializes gRPC server
    GRPCOnly bool
    //HTTPOnly tells orion not to build gRPC server and only initializes HTTP/1.1 server
    HTTPOnly bool
    // HTTPPort is the port to bind for HTTP requests
    HTTPPort string
    // GRPCPost id the port to bind for gRPC requests
    GRPCPort string
    //PprofPort is the port to use for pprof
    PProfport string
    // HotReload when set reloads the service when it receives SIGHUP
    HotReload bool
    //EnableProtoURL adds gRPC generated urls in HTTP handler
    EnableProtoURL bool
    //EnablePrometheus enables prometheus metric for services on path '/metrics' on pprof port
    EnablePrometheus bool
    //EnablePrometheusHistograms enables request histograms for services
    //ref: https://github.com/grpc-ecosystem/go-grpc-prometheus#histograms
    EnablePrometheusHistogram bool
    //HystrixConfig is the configuration options for hystrix
    HystrixConfig HystrixConfig
    //ZipkinConfig is the configuration options for zipkin
    ZipkinConfig ZipkinConfig
    //NewRelicConfig is the configuration options for new relic
    NewRelicConfig NewRelicConfig
    //RollbarToken is the token to be used in rollbar
    RollbarToken string
    //SentryDSN is the token used by sentry for error reporting
    SentryDSN string
    //Env is the environment this service is running in
    Env string
}
```
Config is the configuration used by Orion core

### <a name="BuildDefaultConfig">func</a> [BuildDefaultConfig](./config.go#L79)
``` go
func BuildDefaultConfig(name string) Config
```
BuildDefaultConfig builds a default config object for Orion

## <a name="Decoder">type</a> [Decoder](./types.go#L78)
``` go
type Decoder = handlers.Decoder
```
Decoder is the function type needed for request decoders

## <a name="DefaultServerImpl">type</a> [DefaultServerImpl](./core.go#L70-L86)
``` go
type DefaultServerImpl struct {
    // contains filtered or unexported fields
}
```
DefaultServerImpl provides a default implementation of orion.Server this can be embedded in custom orion.Server implementations

### <a name="DefaultServerImpl.AddDecoder">func</a> (\*DefaultServerImpl) [AddDecoder](./core.go#L161)
``` go
func (d *DefaultServerImpl) AddDecoder(serviceName, method string, decoder handlers.Decoder)
```
AddDecoder is the implementation of handlers.Decodable

### <a name="DefaultServerImpl.AddDefaultDecoder">func</a> (\*DefaultServerImpl) [AddDefaultDecoder](./core.go#L173)
``` go
func (d *DefaultServerImpl) AddDefaultDecoder(serviceName string, decoder Decoder)
```
AddDefaultDecoder is the implementation of handlers.Decodable

### <a name="DefaultServerImpl.AddDefaultEncoder">func</a> (\*DefaultServerImpl) [AddDefaultEncoder](./core.go#L135)
``` go
func (d *DefaultServerImpl) AddDefaultEncoder(serviceName string, encoder Encoder)
```
AddDefaultEncoder is the implementation of handlers.Encodable

### <a name="DefaultServerImpl.AddEncoder">func</a> (\*DefaultServerImpl) [AddEncoder](./core.go#L116)
``` go
func (d *DefaultServerImpl) AddEncoder(serviceName, method string, httpMethod []string, path string, encoder handlers.Encoder)
```
AddEncoder is the implementation of handlers.Encodable

### <a name="DefaultServerImpl.AddHTTPHandler">func</a> (\*DefaultServerImpl) [AddHTTPHandler](./core.go#L143)
``` go
func (d *DefaultServerImpl) AddHTTPHandler(serviceName string, method string, path string, handler handlers.HTTPHandler)
```
AddHTTPHandler is the implementation of handlers.HTTPInterceptor

### <a name="DefaultServerImpl.AddInitializers">func</a> (\*DefaultServerImpl) [AddInitializers](./core.go#L453)
``` go
func (d *DefaultServerImpl) AddInitializers(ins ...Initializer)
```
AddInitializers adds the initializers to orion server

### <a name="DefaultServerImpl.AddMiddleware">func</a> (\*DefaultServerImpl) [AddMiddleware](./core.go#L89)
``` go
func (d *DefaultServerImpl) AddMiddleware(serviceName string, method string, middlewares ...string)
```
AddMiddleware adds middlewares for particular service/method

### <a name="DefaultServerImpl.AddOption">func</a> (\*DefaultServerImpl) [AddOption](./core.go#L181)
``` go
func (d *DefaultServerImpl) AddOption(serviceName, method, option string)
```
AddOption adds a option for the particular service/method

### <a name="DefaultServerImpl.GetConfig">func</a> (\*DefaultServerImpl) [GetConfig](./core.go#L464)
``` go
func (d *DefaultServerImpl) GetConfig() map[string]interface{}
```
GetConfig returns current config as parsed from the file/defaults

### <a name="DefaultServerImpl.GetOrionConfig">func</a> (\*DefaultServerImpl) [GetOrionConfig](./core.go#L194)
``` go
func (d *DefaultServerImpl) GetOrionConfig() Config
```
GetOrionConfig returns current orion config
NOTE: this config can not be modifies

### <a name="DefaultServerImpl.RegisterService">func</a> (\*DefaultServerImpl) [RegisterService](./core.go#L409)
``` go
func (d *DefaultServerImpl) RegisterService(sd *grpc.ServiceDesc, sf interface{}) error
```
RegisterService registers a service from a generated proto file
Note: this is only called from code generated by orion plugin

### <a name="DefaultServerImpl.Start">func</a> (\*DefaultServerImpl) [Start](./core.go#L321)
``` go
func (d *DefaultServerImpl) Start()
```
Start starts the orion server

### <a name="DefaultServerImpl.Stop">func</a> (\*DefaultServerImpl) [Stop](./core.go#L469)
``` go
func (d *DefaultServerImpl) Stop(timeout time.Duration) error
```
Stop stops the server

### <a name="DefaultServerImpl.Wait">func</a> (\*DefaultServerImpl) [Wait](./core.go#L402)
``` go
func (d *DefaultServerImpl) Wait() error
```
Wait waits for all the serving servers to quit

## <a name="Encoder">type</a> [Encoder](./types.go#L75)
``` go
type Encoder = handlers.Encoder
```
Encoder is the function type needed for request encoders

## <a name="FactoryParams">type</a> [FactoryParams](./types.go#L57-L63)
``` go
type FactoryParams struct {
    // ServiceName contains the proto service name
    ServiceName string
    // Version is a counter that is incremented every time a new service object is requested
    // NOTE: version might rollover in long running services
    Version uint64
}
```
FactoryParams are the parameters used by the ServiceFactoryV2

## <a name="HTTPHandler">type</a> [HTTPHandler](./types.go#L81)
``` go
type HTTPHandler = handlers.HTTPHandler
```
HTTPHandler is the http interceptor

## <a name="HystrixConfig">type</a> [HystrixConfig](./config.go#L55-L62)
``` go
type HystrixConfig struct {
    //Port is the port to start hystrix stream handler on
    Port string
    //CommandConfig is configuration for individual commands
    CommandConfig map[string]hystrix.CommandConfig
    //StatsdAddr is the address of the statsd hosts to send hystrix data to
    StatsdAddr string
}
```
HystrixConfig is configuration used by hystrix

### <a name="BuildDefaultHystrixConfig">func</a> [BuildDefaultHystrixConfig](./config.go#L103)
``` go
func BuildDefaultHystrixConfig() HystrixConfig
```
BuildDefaultHystrixConfig builds a default config for hystrix

## <a name="Initializer">type</a> [Initializer](./types.go#L43-L46)
``` go
type Initializer interface {
    Init(svr Server) error
    ReInit(svr Server) error
}
```
Initializer is the interface needed to be implemented by custom initializers

### <a name="ErrorLoggingInitializer">func</a> [ErrorLoggingInitializer](./initializer.go#L45)
``` go
func ErrorLoggingInitializer() Initializer
```
ErrorLoggingInitializer returns a Initializer implementation for error notifier

### <a name="HystrixInitializer">func</a> [HystrixInitializer](./initializer.go#L40)
``` go
func HystrixInitializer() Initializer
```
HystrixInitializer returns a Initializer implementation for Hystrix

### <a name="NewRelicInitializer">func</a> [NewRelicInitializer](./initializer.go#L55)
``` go
func NewRelicInitializer() Initializer
```
NewRelicInitializer returns a Initializer implementation for NewRelic

### <a name="PprofInitializer">func</a> [PprofInitializer](./initializer.go#L65)
``` go
func PprofInitializer() Initializer
```
PprofInitializer returns a Initializer implementation for Pprof

### <a name="PrometheusInitializer">func</a> [PrometheusInitializer](./initializer.go#L60)
``` go
func PrometheusInitializer() Initializer
```
PrometheusInitializer returns a Initializer implementation for Prometheus

### <a name="ZipkinInitializer">func</a> [ZipkinInitializer](./initializer.go#L50)
``` go
func ZipkinInitializer() Initializer
```
ZipkinInitializer returns a Initializer implementation for Zipkin

## <a name="NewRelicConfig">type</a> [NewRelicConfig](./config.go#L71-L76)
``` go
type NewRelicConfig struct {
    APIKey            string
    ServiceName       string
    IncludeAttributes []string
    ExcludeAttributes []string
}
```
NewRelicConfig is the configuration for newrelic

### <a name="BuildDefaultNewRelicConfig">func</a> [BuildDefaultNewRelicConfig](./config.go#L119)
``` go
func BuildDefaultNewRelicConfig() NewRelicConfig
```
BuildDefaultNewRelicConfig builds a default config for newrelic

## <a name="Server">type</a> [Server](./types.go#L25-L40)
``` go
type Server interface {
    //Start starts the orion server, this is non blocking call
    Start()
    //RegisterService registers the service to origin server
    RegisterService(sd *grpc.ServiceDesc, sf interface{}) error
    //Wait waits for the Server loop to exit
    Wait() error
    //Stop stops the Server
    Stop(timeout time.Duration) error
    //GetOrionConfig returns current orion config
    GetOrionConfig() Config
    //GetConfig returns current config as parsed from the file/defaults
    GetConfig() map[string]interface{}
    //AddInitializers adds the initializers to orion server
    AddInitializers(ins ...Initializer)
}
```
Server is the interface that needs to be implemented by any orion server
'DefaultServerImpl' should be enough for most users.

### <a name="GetDefaultServer">func</a> [GetDefaultServer](./core.go#L483)
``` go
func GetDefaultServer(name string) Server
```
GetDefaultServer returns a default server object that can be directly used to start orion server

### <a name="GetDefaultServerWithConfig">func</a> [GetDefaultServerWithConfig](./core.go#L488)
``` go
func GetDefaultServerWithConfig(config Config) Server
```
GetDefaultServerWithConfig returns a default server object that uses provided configuration

## <a name="ServiceFactory">type</a> [ServiceFactory](./types.go#L49-L54)
``` go
type ServiceFactory interface {
    // NewService function receives the server object for which service has to be initialized
    NewService(Server) interface{}
    //DisposeService function disposes the service object
    DisposeService(svc interface{})
}
```
ServiceFactory is the interface that need to be implemented by client that provides with a new service object

## <a name="ServiceFactoryV2">type</a> [ServiceFactoryV2](./types.go#L67-L72)
``` go
type ServiceFactoryV2 interface {
    // NewService function receives the server object for which service has to be initialized
    NewService(svr Server, params FactoryParams) interface{}
    //DisposeService function disposes the service object
    DisposeService(svc interface{}, params FactoryParams)
}
```
ServiceFactoryV2 is the interface that needs to be implemented by client that provides a new service object for multiple services
this allows a single struct to implement multiple services

### <a name="ToServiceFactoryV2">func</a> [ToServiceFactoryV2](./utils.go#L4)
``` go
func ToServiceFactoryV2(sf interface{}) (ServiceFactoryV2, error)
```
ToServiceFactoryV2 converts ServiceFactory to ServiceFactoryV2

## <a name="ZipkinConfig">type</a> [ZipkinConfig](./config.go#L65-L68)
``` go
type ZipkinConfig struct {
    //Addr is the address of the zipkin collector
    Addr string
}
```
ZipkinConfig is the configuration for the zipkin collector

### <a name="BuildDefaultZipkinConfig">func</a> [BuildDefaultZipkinConfig](./config.go#L112)
``` go
func BuildDefaultZipkinConfig() ZipkinConfig
```
BuildDefaultZipkinConfig builds a default config for zipkin

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)