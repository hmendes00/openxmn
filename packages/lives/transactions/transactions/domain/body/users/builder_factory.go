package users

// BuilderFactory represents a UserBuilder factory instance
type BuilderFactory interface {
	Create() Builder
}
