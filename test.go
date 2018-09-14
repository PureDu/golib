package main

import (
	"reflect"
	"unsafe"
	"fmt"
)

func main()  {
	k := reflect.Bool
	k.String()

	fmt.Println(unsafe.Sizeof(uintptr(0)))

	fmt.Println(unsafe.Sizeof(unsafe.Pointer(nil)))
}
