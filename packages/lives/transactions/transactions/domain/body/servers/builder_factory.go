package servers

// BuilderFactory represents a ServerBuilder factory instance
type BuilderFactory interface {
	Create() Builder
}
