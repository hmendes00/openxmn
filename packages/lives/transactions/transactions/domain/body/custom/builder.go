package custom

// Builder represents a Custom builder instance
type Builder interface {
	Create() Builder
	WithCreate(cr Create) Builder
	Now() (Custom, error)
}
