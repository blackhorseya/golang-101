package memento

// Editor 是發起人，負責創建和使用備忘錄來保存和恢復狀態
type Editor struct {
	state string
}

func (e *Editor) SetState(state string) {
	e.state = state
}

func (e *Editor) GetState() string {
	return e.state
}

func (e *Editor) SaveStateToMemento() *Memento {
	return &Memento{state: e.state}
}

func (e *Editor) GetStateFromMemento(memento *Memento) {
	e.state = memento.GetState()
}
