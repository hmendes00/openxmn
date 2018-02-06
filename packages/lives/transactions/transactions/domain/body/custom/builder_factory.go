package custom

// BuilderFactory represents a CustomBuilder factory instance
type BuilderFactory interface {
	Create() Builder
}
