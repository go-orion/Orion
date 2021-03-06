package service

import "github.com/go-orion/Orion/orion"

type svcFactory struct{}

func (s *svcFactory) NewService(orion.Server) interface{} {
	return &svc{}
}

func (s *svcFactory) DisposeService(svc interface{}) {
	//do nothing
}

func GetFactory() orion.ServiceFactory {
	return &svcFactory{}
}
