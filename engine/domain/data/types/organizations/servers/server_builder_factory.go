package servers

// ServerBuilderFactory represents a server builder factory
type ServerBuilderFactory interface {
	Create() ServerBuilder
}
