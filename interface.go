package main

import (
	"fmt"
	"lostfind/playground/zoo"
)

type Pet struct {
	name string
	age  int
}

type MyStringType string

func (m MyStringType) Cry() {
	fmt.Println("I am String")
}

func (p *Pet) Cry() {
	fmt.Println("Nya")
}

func (p *Pet) Walk() {
	fmt.Println("I walk with four feet.")
}

func main() {
	var animal zoo.Animal
	var cat zoo.Cat

	var stringCat zoo.Cat

	myPet := &Pet{
		name: "Nyanko",
		age:  10,
	}

	var myStringType MyStringType

	myStringType = "My String type"

	stringCat = &MyStringType

	cat = myPet
	animal = myPet

	cat.Cry()
	animal.Walk()
	stringCat.Cry()
}
