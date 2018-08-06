package main

import (
	"fmt"
	"reflect"
	"errors"
)

type MyStruct struct {
	name string
}

func (this *MyStruct)GetName() string {
	return this.name
}

func main() {
	s := "this is string"
	ss := errors.New("123")
	fmt.Println(reflect.TypeOf(ss).String())
	fmt.Println("-------------------")

	fmt.Println(reflect.ValueOf(s))
	var x float64 = 3.4
	fmt.Println(reflect.ValueOf(x))
	fmt.Println("-------------------")

	a := new(MyStruct)
	a.name = "yejianfeng"
	typ := reflect.TypeOf(a)

	fmt.Println(typ.NumMethod())
	fmt.Println("-------------------")

	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.TypeOf(a))
	fmt.Println("-------------------")

	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0])

}