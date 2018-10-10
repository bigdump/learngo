package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

func main() {
	cmd := exec.Command(path.Join("/home/test2/hyperchain/go/src/hello/server/test/test"))
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Over!")
	<-time.After(12 * time.Second)
	fmt.Println("Over again!!")

}
