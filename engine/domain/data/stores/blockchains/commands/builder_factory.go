package commands

// BuilderFactory represents a stored commands builder factory
type BuilderFactory interface {
	Create() Builder
}
