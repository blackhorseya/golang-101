package command

import "fmt"

// Light 是接收者，執行具體操作
type Light struct{}

func (l *Light) On() {
	fmt.Println("The light is on")
}

func (l *Light) Off() {
	fmt.Println("The light is off")
}
