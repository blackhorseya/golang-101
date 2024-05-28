package state

import "fmt"

// OnState 是具體狀態，表示燈是打開的狀態
type OnState struct{}

func (s *OnState) PressSwitch(light *Light) {
	fmt.Println("Turning off the light.")
	light.SetState(&OffState{})
}

// OffState 是具體狀態，表示燈是關閉的狀態
type OffState struct{}

func (s *OffState) PressSwitch(light *Light) {
	fmt.Println("Turning on the light.")
	light.SetState(&OnState{})
}
