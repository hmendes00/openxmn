package commands

// BuilderFactory represents a commands builder factory
type BuilderFactory interface {
	Create() Builder
}
