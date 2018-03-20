package chunks

// BuilderFactory represents a ChunksBuilder factory
type BuilderFactory interface {
	Create() Builder
}
