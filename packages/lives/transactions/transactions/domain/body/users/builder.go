package users

// Builder represents a User builder instance
type Builder interface {
	Create() Builder
	WithSave(cr Save) Builder
	WithDelete(del Delete) Builder
	Now() (User, error)
}
