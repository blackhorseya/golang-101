package main

import (
	"fmt"
)

// MarketData is a context that holds the necessary data for market data.
type MarketData struct {
	// Symbol is the symbol of the market data
	Symbol string `json:"symbol,omitempty"`

	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Low    float64 `json:"low"`
	Volume float64 `json:"volume"`
}

// GetIndicator returns the value of the specified indicator.
func (x *MarketData) GetIndicator(indicator string) float64 {
	switch indicator {
	case "open":
		return x.Open
	case "high":
		return x.High
	case "close", "price":
		return x.Close
	case "low":
		return x.Low
	case "volume":
		return x.Volume
	}

	return 0
}

// StrategyParameters that encapsulates the parameters of a trading strategy.
type StrategyParameters struct {
	// EntryConditions define the conditions under which the strategy will enter a trade
	EntryConditions []Condition `json:"entry_conditions,omitempty"`

	// ExitConditions define the conditions under which the strategy will exit a trade
	ExitConditions []Condition `json:"exit_conditions,omitempty"`
}

// NewStrategy creates a new strategy with the given entry and exit conditions.
func NewStrategy(entryConditions []Condition, exitConditions []Condition) *StrategyParameters {
	return &StrategyParameters{
		EntryConditions: entryConditions,
		ExitConditions:  exitConditions,
	}
}

// GetEntryExpression returns the entry expression of the strategy.
func (x *StrategyParameters) GetEntryExpression() Expression {
	return x.getExpression(x.EntryConditions)
}

func (x *StrategyParameters) GetExitExpression() Expression {
	return x.getExpression(x.ExitConditions)
}

func (x *StrategyParameters) getExpression(conditions []Condition) Expression {
	if len(conditions) == 0 {
		return &NilExpression{}
	}

	if len(conditions) == 1 {
		return &conditions[0]
	}

	return &AndExpression{
		Left:  &conditions[0],
		Right: x.getExpression(conditions[1:]),
	}
}

// Condition is a value object that represents a single condition used in a trading strategy.
type Condition struct {
	// Symbol is the trading symbol (e.g., stock ticker, currency pair) for which the condition applies
	Symbol string `json:"symbol,omitempty"`

	// Indicator is the market indicator used in the condition (e.g., moving average, RSI)
	Indicator string `json:"indicator,omitempty"`

	// Operator is the comparison operator used in the condition (e.g., >, <, ==)
	Operator string `json:"operator,omitempty"`

	// Value is the value to compare the indicator against
	Value float64 `json:"value,omitempty"`
}

func (c *Condition) Interpret(context map[string]*MarketData) bool {
	data, ok := context[c.Symbol]
	if !ok {
		return false
	}

	indicator := data.GetIndicator(c.Indicator)

	switch c.Operator {
	case ">":
		return indicator > c.Value
	case "<":
		return indicator < c.Value
	case "==":
		return indicator == c.Value
	}

	return false
}

func main() {
	// Create a context
	data1 := &MarketData{
		Symbol: "AAPL",
		Open:   100.0,
		High:   110.0,
		Close:  105.0,
		Low:    95.0,
		Volume: 1000000.0,
	}

	data2 := &MarketData{
		Symbol: "GOOG",
		Open:   1000.0,
		High:   1100.0,
		Close:  1050.0,
		Low:    950.0,
		Volume: 10000000.0,
	}

	context := map[string]*MarketData{
		"AAPL": data1,
		"GOOG": data2,
	}

	strategy := NewStrategy([]Condition{
		{
			Symbol:    "AAPL",
			Indicator: "close",
			Operator:  ">",
			Value:     100.0,
		},
		{
			Symbol:    "GOOG",
			Indicator: "volume",
			Operator:  ">",
			Value:     5000000.0,
		},
	}, []Condition{})

	entryExpression := strategy.GetEntryExpression()
	if entryExpression.Interpret(context) {
		fmt.Println("Enter trade")
	} else {
		fmt.Println("Do not enter trade")
	}

	exitExpression := strategy.GetExitExpression()
	if exitExpression.Interpret(context) {
		fmt.Println("Exit trade")
	} else {
		fmt.Println("Do not exit trade")
	}
}
