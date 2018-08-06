package main

import "fmt"

func main()  {
	fmt.Print(a())
}

func a() string {
	defer func() string{
		return "234"
	}()
	return "123"
}

