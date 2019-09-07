package main

import (
	"fmt"
	"time"
	"unsafe"
)

type User struct {
	name  string
	birth time.Time
}

func main() {
	intVal := 5
	intPointer := &intVal

	fmt.Println(intVal)      // 5
	fmt.Println(*intPointer) // 5

	*intPointer = 10 // intVal = 10 するのと同じ

	fmt.Println(intVal)      // 10
	fmt.Println(*intPointer) // 10

	var userPointer *User
	var boolPointer *bool

	fmt.Println(unsafe.Sizeof(intPointer))  // 8
	fmt.Println(unsafe.Sizeof(userPointer)) // 8
	fmt.Println(unsafe.Sizeof(boolPointer)) // 8
}
