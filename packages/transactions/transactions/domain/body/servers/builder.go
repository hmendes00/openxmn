package servers

// Builder represents a Server builder instance
type Builder interface {
	Create() Builder
	WithCreate(cr Create) Builder
	WithDelete(del Delete) Builder
	Now() (Server, error)
}
