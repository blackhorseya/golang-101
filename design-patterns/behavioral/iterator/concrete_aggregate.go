package iterator

// NameCollection 是具體的聚合，實現了 Aggregate 接口
type NameCollection struct {
	names []string
}

func (n *NameCollection) CreateIterator() Iterator {
	return &NameIterator{
		names: n.names,
		index: 0,
	}
}

func NewNameCollection(names []string) *NameCollection {
	return &NameCollection{names: names}
}
