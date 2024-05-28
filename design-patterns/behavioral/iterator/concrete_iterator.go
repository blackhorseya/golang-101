package iterator

// NameIterator 是具體的迭代器，實現了 Iterator 接口
type NameIterator struct {
	names []string
	index int
}

func (n *NameIterator) HasNext() bool {
	return n.index < len(n.names)
}

func (n *NameIterator) Next() interface{} {
	if n.HasNext() {
		name := n.names[n.index]
		n.index++
		return name
	}
	return nil
}
