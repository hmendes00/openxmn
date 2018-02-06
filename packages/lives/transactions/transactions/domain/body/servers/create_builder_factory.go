package servers

// CreateBuilderFactory represents the builder factory of a create server transaction
type CreateBuilderFactory interface {
	Create() CreateBuilder
}
