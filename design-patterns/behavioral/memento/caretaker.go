package memento

// Caretaker 是負責人，保存備忘錄，但不能操作或檢查備忘錄的內容
type Caretaker struct {
	mementoList []*Memento
}

func (c *Caretaker) Add(memento *Memento) {
	c.mementoList = append(c.mementoList, memento)
}

func (c *Caretaker) Get(index int) *Memento {
	if index < len(c.mementoList) {
		return c.mementoList[index]
	}
	return nil
}
