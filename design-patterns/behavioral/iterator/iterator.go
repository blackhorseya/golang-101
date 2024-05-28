package iterator

// Iterator 定義了迭代器的接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}
