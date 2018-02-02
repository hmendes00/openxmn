package users

// CreateBuilderFactory represents the builder factory of a create user transaction
type CreateBuilderFactory interface {
	Create() CreateBuilder
}
