package mediator

import (
	"testing"
)

func TestMediatorPattern(t *testing.T) {
	mediator := &ChatMediator{}

	user1 := NewUser("User1", mediator)
	user2 := NewUser("User2", mediator)
	user3 := NewUser("User3", mediator)

	mediator.AddColleague(user1)
	mediator.AddColleague(user2)
	mediator.AddColleague(user3)

	user1.SendMessage("Hello, everyone!")
	user2.SendMessage("Hi, User1!")

	// 這裡可以添加更多的測試邏輯來驗證消息是否正確傳遞
}
