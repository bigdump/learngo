package main

import (
	"fmt"
	"github.com/op/go-logging"
)

type impl1 struct {
	gas int64
	vm string
	log logging.Logger
}

func newImpl1() *impl1 {
	return &impl1{
		gas: 100,
	}
}

func (imp *impl1) Create(namespace string, data []byte) {
	fmt.Println("I am in imp1 create")
}

func (imp *impl1) Call(address string, input []byte) []byte {
	fmt.Println("I am in imp1 call")
	return []byte{}
}
