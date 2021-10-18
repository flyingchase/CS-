package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	D bool
	C string
	B []int32
	A int32
}

func main() {
	var u User
	fmt.Println(unsafe.Sizeof(u))
}
