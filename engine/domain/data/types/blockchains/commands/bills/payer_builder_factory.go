package bills

// PayerBuilderFactory represents a payer builder factory
type PayerBuilderFactory interface {
	Create() PayerBuilder
}
