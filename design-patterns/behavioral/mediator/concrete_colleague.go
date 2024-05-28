package mediator

import "fmt"

// User 是具體同事，通過中介者與其他同事通信
type User struct {
	name     string
	mediator Mediator
}

func NewUser(name string, mediator Mediator) *User {
	return &User{name: name, mediator: mediator}
}

func (u *User) ReceiveMessage(message string) {
	fmt.Printf("%s received: %s\n", u.name, message)
}

func (u *User) SendMessage(message string) {
	fmt.Printf("%s sends: %s\n", u.name, message)
	u.mediator.SendMessage(message, u)
}
