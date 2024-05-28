package interpreter

import "strconv"

// Number 終結符表達式，用於解析數字
type Number struct {
	value int
}

func (n *Number) Interpret() int {
	return n.value
}

func NewNumber(value string) (*Number, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}
	return &Number{value: v}, nil
}
