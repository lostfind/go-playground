package main

import (
	"fmt"
	"lostfind/playground/di/light"
)

type RoomNotDI struct {
	LightOne light.Incandescent
	LightTwo light.LedLight
}

func (r *RoomNotDI) SwitchOnOne() {
	fmt.Println("1番照明:", r.LightOne.LightUp())
}
func (r *RoomNotDI) SwitchOnTwo() {
	fmt.Println("2番照明:", r.LightTwo.LightUp())
}

func main() {
	myRoom := new(RoomNotDI)
	myRoom.SwitchOnOne()
	myRoom.SwitchOnTwo()
}
