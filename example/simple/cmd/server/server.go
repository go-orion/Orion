package main

import (
	"fmt"

	"github.com/go-orion/Orion/example/simple/service"
	proto "github.com/go-orion/Orion/example/simple/simple_proto"
	"github.com/go-orion/Orion/orion"
)

type svcFactory struct {
}

func (s *svcFactory) NewService(svr orion.Server) interface{} {
	return service.GetService()
}

func (s *svcFactory) DisposeService(svc interface{}) {
	fmt.Println("disposing", svc)
}

func main() {
	server := orion.GetDefaultServer("SimpleService")
	proto.RegisterSimpleServiceOrionServer(&svcFactory{}, server)
	server.Start()
	server.Wait()
}
