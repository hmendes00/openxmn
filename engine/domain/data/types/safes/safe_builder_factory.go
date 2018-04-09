package safes

// SafeBuilderFactory represents a safe builder factory
type SafeBuilderFactory interface {
	Create() SafeBuilder
}
