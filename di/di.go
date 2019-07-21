package main

import (
	"fmt"
	"lostfind/playground/di/light"
)

type LightSocket interface {
	LightUp() string
}

type RoomDI struct {
	LightOne LightSocket
	LightTwo LightSocket
}

func NewRoom(lightOne, lightTwo LightSocket) *RoomDI {
	room := &RoomDI{
		LightOne: lightOne,
		LightTwo: lightTwo,
	}
	return room
}

func (r *RoomDI) SwitchOnOne() {
	fmt.Println("1番照明:", r.LightOne.LightUp())
}
func (r *RoomDI) SwitchOnTwo() {
	fmt.Println("2番照明:", r.LightTwo.LightUp())
}

func main() {
	myRoom := NewRoom(new(light.Incandescent), new(light.LedLight)) // ここがDI

	myRoom.SwitchOnOne()
	myRoom.SwitchOnTwo()

	myRoom.LightOne = new(light.Incandescent)
	myRoom.LightTwo = new(light.LedLight)

}
