package main

import (
	"archive/zip"
	"fmt"
)

func main() {
	fmt.Print(a())
}

func a() string {
	reader, err := zip.OpenReader("/home/june/hyperchain/litesdk/src/test/resources/hvm-jar/hvmbasic-1.0.0-student.jar")
	if err != nil {
		fmt.Println("f")
	}

	for _, zipf := range reader.File {
		fmt.Println(zipf.Name)
	}

	defer func() string {
		return "234"
	}()
	return "123"
}
