package main

import (
	"fmt"
	"net/http"
	"runtime/pprof"
	"time"
)

//var quit chan struct{} = make(chan struct{})

func f() {
	a := 1231
	b := -1
	a = a + b
}

func sum(a, b int) int {
	return a + b
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func main() {
	go func() {
		http.HandleFunc("/", handler)
		http.ListenAndServe(":11181", nil)
	}()

	for i := 0; i < 100000000000000; i++ {
		if i % 10000000000 == 0 {
			fmt.Println("--------")
		}
		go f()
	}
	fmt.Println("+++++++++")
	time.Sleep(10000 * time.Second)
}
