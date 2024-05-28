package iterator

// Aggregate 定義了創建迭代器的接口
type Aggregate interface {
	CreateIterator() Iterator
}
