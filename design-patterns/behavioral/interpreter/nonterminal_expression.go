package interpreter

// Add 非終結符表達式，用於解析加法表達式
type Add struct {
	left  Expression
	right Expression
}

func (a *Add) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

// Subtract 非終結符表達式，用於解析減法表達式
type Subtract struct {
	left  Expression
	right Expression
}

func (s *Subtract) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}
