package servers

// PriceBuilderFactory represents a server price builder factory
type PriceBuilderFactory interface {
	Create() PriceBuilder
}
