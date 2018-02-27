package domain

// ChunksBuilderFactory represents a ChunksBuilder factory
type ChunksBuilderFactory interface {
	Create() ChunksBuilder
}
