package users

// Builder represents a User builder instance
type Builder interface {
	Create() Builder
	WithCreate(cr Create) Builder
	WithDelete(del Delete) Builder
	WithUpdate(up Update) Builder
	Now() (User, error)
}
