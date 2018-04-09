package bills

// BillBuilderFactory represents a bill builder factory
type BillBuilderFactory interface {
	Create() BillBuilder
}
