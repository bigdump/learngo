package main

import (
	"fmt"
	"sync"
)

var p sync.Pool

func main() {

	a := p.Get()
	if a != nil {
		fmt.Println(a.([]byte)[:10])
	} else {
		f := make([]byte, 10)
		p.Put(f)
		fmt.Println(f)

		ag := make([]byte, 5)
		ag[0] = 31
		ag[3] = 4
		fmt.Println(ag)
		copy(f[9:], ag)
		fmt.Println(f)
		fmt.Println(ag)
	}
	fmt.Println()
}
