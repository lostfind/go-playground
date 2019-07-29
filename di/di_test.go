package main

import "testing"

type MockLight struct{}

func (*MockLight) LightUp() string {
	return "ぴーぴー音を鳴らす"
}

func TestSwitchOnOne(t *testing.T) {
	room := NewRoom(new(MockLight), new(MockLight))

	room.SwitchOnOne()
}
