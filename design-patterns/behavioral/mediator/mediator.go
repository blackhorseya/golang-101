package mediator

// Mediator 定義了中介者的接口
type Mediator interface {
	SendMessage(message string, colleague Colleague)
	AddColleague(colleague Colleague)
}
