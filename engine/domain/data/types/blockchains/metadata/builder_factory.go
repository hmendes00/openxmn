package metadata

// BuilderFactory represents a Builder factory
type BuilderFactory interface {
	Create() Builder
}
