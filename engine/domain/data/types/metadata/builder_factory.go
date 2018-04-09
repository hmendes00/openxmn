package metadata

// BuilderFactory represents a metadata builder factory
type BuilderFactory interface {
	Create() Builder
}
