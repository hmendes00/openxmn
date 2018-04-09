package bills

// PayeeBuilderFactory represents a payee builder factory
type PayeeBuilderFactory interface {
	Create() PayeeBuilder
}
