package main

import (
	"fmt"
	"os"
)

type Test struct {
	ID      int
	Name    string
	Address string
}

type node1 struct {
	x int
	y int
}

type node2 struct {
	node1
	z int
	x int
}

func Testp(id int, test *Test) Test {
	test.ID = id
	return *test
}

func main() {
	a := "hello, world\n"
	fmt.Printf(a)

	sep := "       "
	for index, value := range os.Args[:] {
		fmt.Print(index)
		fmt.Println(sep + value)
	}
	fmt.Println("Done.")

	tt := Test{1232, "1231", "12341"}

	fmt.Println(Testp(123, &tt).ID)
	fmt.Println(tt.ID)

	w := node2{
		node1: node1{x: 6, y: 6},
		z:     5,
	}

	w.x = 3333
	w.y = 3333
	fmt.Println(w.x)
	fmt.Println(w.node1.x)

}
