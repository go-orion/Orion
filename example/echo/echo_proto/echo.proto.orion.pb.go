// Code generated by protoc-gen-orion. DO NOT EDIT.
// source: echo.proto

package echo_proto

import (
	orion "github.com/carousell/Orion/orion"
)

func RegisterEchoServiceOrionServer(srv orion.ServiceFactory, orionServer orion.Server) {
	orionServer.RegisterService(&_EchoService_serviceDesc, srv)
}

func RegisterEchoServiceUpperEncoder(svr orion.Server, encoder orion.Encoder) {
	orion.RegisterEncoder(svr, "EchoService", "Upper","GET", "/api/1.0/upper/{msg}", encoder)
}
