package domain

// UserBuilderFactory represents a UserBuilder factory
type UserBuilderFactory interface {
	Create() UserBuilder
}
