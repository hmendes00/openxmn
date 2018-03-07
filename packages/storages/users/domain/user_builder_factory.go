package domain

// UserBuilderFactory represents a stored user builder factory
type UserBuilderFactory interface {
	Create() UserBuilder
}
