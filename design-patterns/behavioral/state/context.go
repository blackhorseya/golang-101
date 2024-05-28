package state

// Light 是上下文類，持有當前的狀態對象
type Light struct {
	state State
}

func NewLight() *Light {
	return &Light{state: &OffState{}} // 初始狀態為關閉
}

func (l *Light) SetState(state State) {
	l.state = state
}

func (l *Light) PressSwitch() {
	l.state.PressSwitch(l)
}
