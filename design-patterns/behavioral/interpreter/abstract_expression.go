package interpreter

// Expression 定義了抽象表達式接口
type Expression interface {
	Interpret() int
}
