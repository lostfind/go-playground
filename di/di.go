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

	lightOne := new(light.Incandescent) // 注入するオブジェクト
	lightTwo := new(light.LedLight)     // 注入するオブジェクト

	myRoom := NewRoom(lightOne, lightTwo) // ここがDI

	myRoom.SwitchOnOne()
	myRoom.SwitchOnTwo()
}
