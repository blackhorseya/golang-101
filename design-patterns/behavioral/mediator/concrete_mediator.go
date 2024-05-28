package mediator

// ChatMediator 是具體的中介者，協調同事對象之間的交互
type ChatMediator struct {
	colleagues []Colleague
}

func (m *ChatMediator) AddColleague(colleague Colleague) {
	m.colleagues = append(m.colleagues, colleague)
}

func (m *ChatMediator) SendMessage(message string, sender Colleague) {
	for _, colleague := range m.colleagues {
		if colleague != sender {
			colleague.ReceiveMessage(message)
		}
	}
}
