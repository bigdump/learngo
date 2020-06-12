package main

import (
	"fmt"
	"io"
	"os"
)

const FileSize = 1024

func main() {
	fd, err := os.OpenFile("file.txt", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
	st, err := fd.Stat()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(st.Size())
	n := st.Size() / FileSize
	if st.Size()%FileSize > 0 {
		n++
	}

	for i := int64(0); i < n; i++ {
		bt := make([]byte, FileSize)
		offset := i * FileSize
		r := io.NewSectionReader(fd, offset, FileSize)
		n, err := r.Read(bt)
		if err!= nil {
			fmt.Println(err.Error())
			fmt.Println("====================")
		}
		fmt.Printf("Read file size : %v\n", n)
		fmt.Printf("Read file size of bytes: %v\n, the bytes is : %v", len(bt), bt)
		fmt.Println("")
	}
}
