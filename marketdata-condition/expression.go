package main

// Expression is an interface that represents a mathematical expression.
type Expression interface {
	// Interpret interprets the expression in the given context.
	Interpret(context map[string]*MarketData) bool
}

// NilExpression is a concrete expression that represents a nil value.
type NilExpression struct {
}

func (x *NilExpression) Interpret(context map[string]*MarketData) bool {
	return false
}

// AndExpression is a concrete expression that represents a logical AND operation.
type AndExpression struct {
	Left  Expression
	Right Expression
}

func (ae *AndExpression) Interpret(context map[string]*MarketData) bool {
	return ae.Left.Interpret(context) && ae.Right.Interpret(context)
}

// OrExpression is a concrete expression that represents a logical OR operation.
type OrExpression struct {
	Left  Expression
	Right Expression
}

func (oe *OrExpression) Interpret(context map[string]*MarketData) bool {
	return oe.Left.Interpret(context) || oe.Right.Interpret(context)
}
