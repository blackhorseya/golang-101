package visitor

// Shape 定義了形狀的接口，包含接受訪問者的方法
type Shape interface {
	Accept(visitor Visitor)
}
