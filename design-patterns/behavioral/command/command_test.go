package command

import (
	"testing"
)

func TestLightCommands(t *testing.T) {
	light := &Light{}

	lightOnCommand := NewLightOnCommand(light)
	lightOffCommand := NewLightOffCommand(light)

	remote := &RemoteControl{}

	// 測試打開燈
	remote.SetCommand(lightOnCommand)
	remote.PressButton()

	// 測試關閉燈
	remote.SetCommand(lightOffCommand)
	remote.PressButton()
}
