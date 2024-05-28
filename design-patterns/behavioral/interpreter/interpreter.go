package interpreter

import (
	"fmt"
	"strings"
)

// Parser 解析器，用於解析表達式
type Parser struct {
	tokens []string
}

func NewParser(expression string) *Parser {
	tokens := strings.Fields(expression)
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() (Expression, error) {
	if len(p.tokens) == 0 {
		return nil, fmt.Errorf("no tokens to parse")
	}

	var stack []Expression

	for _, token := range p.tokens {
		switch token {
		case "+":
			if len(stack) < 2 {
				return nil, fmt.Errorf("invalid expression")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, &Add{left: left, right: right})
		case "-":
			if len(stack) < 2 {
				return nil, fmt.Errorf("invalid expression")
			}
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, &Subtract{left: left, right: right})
		default:
			number, err := NewNumber(token)
			if err != nil {
				return nil, err
			}
			stack = append(stack, number)
		}
	}

	if len(stack) != 1 {
		return nil, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}
