package orion

import (
	"github.com/go-orion/Orion/orion/handlers"
)

//RegisterEncoder allows for registering an HTTP request encoder to arbitrary urls
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterEncoder(svr Server, serviceName, method, httpMethod, path string, encoder Encoder) {
	if e, ok := svr.(handlers.Encodeable); ok {
		e.AddEncoder(serviceName, method, []string{httpMethod}, path, encoder)
	}
}

//RegisterEncoders allows for registering an HTTP request encoder to arbitrary urls
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterEncoders(svr Server, serviceName, method string, httpMethod []string, path string, encoder Encoder) {
	if e, ok := svr.(handlers.Encodeable); ok {
		e.AddEncoder(serviceName, method, httpMethod, path, encoder)
	}
}

//RegisterDefaultEncoder allows for registering an HTTP request encoder to arbitrary urls for the entire service
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterDefaultEncoder(svr Server, serviceName string, encoder Encoder) {
	if e, ok := svr.(handlers.Encodeable); ok {
		e.AddDefaultEncoder(serviceName, encoder)
	}
}

//RegisterDefaultDecoder allows for registering an HTTP request decoder to arbitrary urls for the entire service
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterDefaultDecoder(svr Server, serviceName string, decoder Decoder) {
	if e, ok := svr.(handlers.Decodable); ok {
		e.AddDefaultDecoder(serviceName, decoder)
	}
}

//RegisterDecoder allows for registering an HTTP request decoder to a method
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterDecoder(svr Server, serviceName, method string, decoder Decoder) {
	if e, ok := svr.(handlers.Decodable); ok {
		e.AddDecoder(serviceName, method, decoder)
	}
}

//RegisterHandler allows registering an HTTP handler for a given path
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterHandler(svr Server, serviceName, method string, path string, handler HTTPHandler) {
	if e, ok := svr.(handlers.HTTPInterceptor); ok {
		e.AddHTTPHandler(serviceName, method, path, handler)
	}
}

//RegisterMethodOption allows for registering an handler option to a particular method
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterMethodOption(svr Server, serviceName, method, option string) {
	if e, ok := svr.(handlers.Optionable); ok {
		e.AddOption(serviceName, method, option)
	}
}

//RegisterMiddleware allows for registering  middlewares to a particular method
//Note: this is normally called from protoc-gen-orion autogenerated files
func RegisterMiddleware(svr Server, serviceName, method string, middleware ...string) {
	if e, ok := svr.(handlers.Middlewareable); ok {
		e.AddMiddleware(serviceName, method, middleware...)
	}
}
