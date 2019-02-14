// Code generated by protoc-gen-orion. DO NOT EDIT.
// source: echo.proto

package echo_proto

import (
	orion "github.com/go-orion/Orion/orion"
)

// If you see error please update your orion-protoc-gen by running 'go get -u github.com/go-orion/Orion/protoc-gen-orion'
var _ = orion.ProtoGenVersion1_0

// Encoders

// RegisterEchoServiceUpperEncoder registers the encoder for Upper method in EchoService
// it registers HTTP  path /api/1.0/upper/{msg} with "GET", "POST", "OPTIONS" methods
func RegisterEchoServiceUpperEncoder(svr orion.Server, encoder orion.Encoder) {
	orion.RegisterEncoders(svr, "EchoService", "Upper", []string{"GET", "POST", "OPTIONS"}, "/api/1.0/upper/{msg}", encoder)
}

// RegisterEchoServiceUpperProxyEncoder registers the encoder for UpperProxy method in EchoService
// it registers HTTP with "POST", "PUT" methods
func RegisterEchoServiceUpperProxyEncoder(svr orion.Server, encoder orion.Encoder) {
	orion.RegisterEncoders(svr, "EchoService", "UpperProxy", []string{"POST", "PUT"}, "", encoder)
}

// Handlers

// RegisterEchoServiceUpperHandler registers the handler for Upper method in EchoService
func RegisterEchoServiceUpperHandler(svr orion.Server, handler orion.HTTPHandler) {
	orion.RegisterHandler(svr, "EchoService", "Upper", "/api/1.0/upper/{msg}", handler)
}

// RegisterEchoServiceUpperProxyHandler registers the handler for UpperProxy method in EchoService
func RegisterEchoServiceUpperProxyHandler(svr orion.Server, handler orion.HTTPHandler) {
	orion.RegisterHandler(svr, "EchoService", "UpperProxy", "", handler)
}

// Decoders

// RegisterEchoServiceUpperDecoder registers the decoder for Upper method in EchoService
func RegisterEchoServiceUpperDecoder(svr orion.Server, decoder orion.Decoder) {
	orion.RegisterDecoder(svr, "EchoService", "Upper", decoder)
}

// RegisterEchoServiceUpperProxyDecoder registers the decoder for UpperProxy method in EchoService
func RegisterEchoServiceUpperProxyDecoder(svr orion.Server, decoder orion.Decoder) {
	orion.RegisterDecoder(svr, "EchoService", "UpperProxy", decoder)
}

// RegisterEchoServiceOrionServer registers EchoService to Orion server
func RegisterEchoServiceOrionServer(srv orion.ServiceFactory, orionServer orion.Server) {
	orionServer.RegisterService(&_EchoService_serviceDesc, srv)

	RegisterEchoServiceUpperEncoder(orionServer, nil)
	RegisterEchoServiceUpperProxyEncoder(orionServer, nil)
}
