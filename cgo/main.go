package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lhvm

#include <stdio.h>
#include <stdlib.h>
#include "hvm.h"

void print(char *str) {
   printf("%s\n", str);
}
*/
import "C"

import "unsafe"

func main() {
	s := "Hello Cgo"
	cs := C.CString(s)
	C.print(cs)
	C.free(unsafe.Pointer(cs))
	C.hvm_invoke()
	C.hvm_deploy()
}
