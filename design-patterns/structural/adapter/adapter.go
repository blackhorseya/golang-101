package adapter

// Shape 定義形狀的接口
type Shape interface {
	Draw() string
}

// LegacyCircle 是舊版的圓形類，具有不兼容的新方法
type LegacyCircle struct{}

func (lc *LegacyCircle) OldDraw() string {
	return "Drawing a circle using the old method"
}

// CircleAdapter 是適配器，將 LegacyCircle 的接口轉換為 Shape 接口
type CircleAdapter struct {
	LegacyCircle *LegacyCircle
}

func (ca *CircleAdapter) Draw() string {
	// 使用 LegacyCircle 的 OldDraw 方法
	return ca.LegacyCircle.OldDraw()
}
