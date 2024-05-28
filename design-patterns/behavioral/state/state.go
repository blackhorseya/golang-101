package state

// State 定義了狀態的接口
type State interface {
	PressSwitch(light *Light)
}
