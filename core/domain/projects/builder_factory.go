package projects

// BuilderFactory represents a projects builder factory
type BuilderFactory interface {
	Create() Builder
}
