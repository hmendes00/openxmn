package custom

// CreateBuilderFactory represents the builder factory of a create custom transaction
type CreateBuilderFactory interface {
	Create() CreateBuilder
}
