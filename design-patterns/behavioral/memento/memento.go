package memento

// Memento 是備忘錄，負責存儲發起人的內部狀態
type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

func (m *Memento) SetState(state string) {
	m.state = state
}
