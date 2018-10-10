package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	setupDebug()
	start := time.Now().UnixNano()
	b := make(map[string][]byte)
	for i := 0; i < 1000000; i++ {
		time.After(time.Microsecond)
		b[string(i)] = make([]byte, 900, 3000)
		c := new(big.Int)
		c.MarshalJSON()
	}
	for i := 0; i < 1000000; i++ {
		time.After(time.Microsecond)
		b[string(i)] = make([]byte, 900, 3000)
		c := new(big.Int)
		c.MarshalJSON()
	}
	for i := 0; i < 1000000; i++ {
		time.After(time.Microsecond)
		b[string(i)] = make([]byte, 900, 3000)
		c := new(big.Int)
		c.MarshalJSON()
	}
	end := time.Now().UnixNano()
	fmt.Println(b[string(10)])
	malloc()
	fmt.Println("I am done")
	fmt.Printf("Use time : %v ms \n", (end-start)/1000/1000)
}

type Node struct {
	Bg *big.Int
	s  string
}

func malloc() {
	for i := 0; i < 10; i++ {
		c := &Node{
			Bg: big.NewInt(12),
			s:  "12345",
		}
		c.Bg = big.NewInt(13)
	}
}
