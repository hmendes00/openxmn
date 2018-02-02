package body

// BuilderFactory represents the BodyBuilder factory
type BuilderFactory interface {
	Create() Builder
}
