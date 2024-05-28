package mediator

// Colleague 定義了同事的接口
type Colleague interface {
	ReceiveMessage(message string)
	SendMessage(message string)
}
