package projects

// BuilderFactory represents a servers builder factory
type BuilderFactory interface {
	Create() Builder
}
