package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("I am starting")
	<-time.After(10 * time.Second)
	fmt.Println("I am finished !")
}
