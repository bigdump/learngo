package main

import "github.com/op/go-logging"

type impl2 struct {
	gas int64
	vm string
	log logging.Logger
}

func (imp *impl2) Create(namespace string, data []byte) {
	panic("implement me")
}

func (imp *impl2) Call(address string, input []byte) []byte {
	panic("implement me")
}
