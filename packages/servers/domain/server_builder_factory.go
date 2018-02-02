package domain

// ServerBuilderFactory represents the ServerBuilder factory
type ServerBuilderFactory interface {
	Create() ServerBuilder
}
