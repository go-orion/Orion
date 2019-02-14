package main

import (
	"github.com/go-orion/Orion/example/stringsvc2/service"
	proto "github.com/go-orion/Orion/example/stringsvc2/stringproto"
	"github.com/go-orion/Orion/orion"
)

func main() {
	server := orion.GetDefaultServer("StringService")
	proto.RegisterStringServiceOrionServer(service.GetFactory(), server)
	server.Start()
	server.Wait()
}
