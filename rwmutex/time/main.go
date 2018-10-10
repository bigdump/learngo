package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {

	//noLock()
	//abc := fmt.Sprintf("decompress source contract failed %v","admin")
	//fmt.Println(abc)
	//
	//before := time.Now().UnixNano()
	//for i := 0; i < 1000; i++ {
	//	noLock()
	//}
	//after := time.Now().UnixNano()
	//for i := 0; i < 1000; i++ {
	//	Lock()
	//}
	//after2 := time.Now().UnixNano()
	//
	//fmt.Printf("noLock : %v ns\n", (after - before)/1000)
	//fmt.Printf("Lock : %v ns\n", (after2 - after)/1000)
	//killProcess()
	//
	//<- time.After(time.Second * 500)
	getfiles()

}

func getfiles () {
	cmd := exec.Command("ls", "-a", "/home/june")

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。 这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	cmd.Run()
	fmt.Println(out.String())
}


func killProcess() {
	er := make(chan string)
	go func() {
		err := prs()
		er <- err.Error()
	}()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Time out !!!!!!!!!!!!!")
	case rs := <-er:
		fmt.Println("not time out")
		fmt.Println(rs)
	}
}

func prs() error {
	for i := 0; i < 10000; i++ {
		println("I am alive!")
		<- time.After(time.Millisecond * 500)
	}
	return errors.New("123")
}

func noLock() string{
	a := "123"
	fmt.Print(a)
	defer fmt.Println("123")
	return a
	return "12345"

}

func Lock() {
	lock.Lock()
	defer lock.Unlock()
}
