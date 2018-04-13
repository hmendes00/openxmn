package safes

// AmountBuilderFactory represents an amount builder factory
type AmountBuilderFactory interface {
	Create() AmountBuilder
}
