package  main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Name string
	Age int
}

func main() {
	u:=new(User)
	name:=(*string)(unsafe.Pointer(u))
	*name="Hello"

	age:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u))+unsafe.Offsetof(u.Age)))
	*age=17
	fmt.Println(u)
}