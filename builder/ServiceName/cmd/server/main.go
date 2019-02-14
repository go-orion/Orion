package main

import (
	proto "github.com/go-orion/Orion/builder/ServiceName/ServiceName_proto"
	"github.com/go-orion/Orion/builder/ServiceName/service"
	"github.com/go-orion/Orion/orion"
)

func main() {
	server := orion.GetDefaultServer("EchoService")
	proto.RegisterServiceNameOrionServer(service.GetServiceFactory(), server)
	service.RegisterOptionals(server)
	server.Start()
	server.Wait()
}
