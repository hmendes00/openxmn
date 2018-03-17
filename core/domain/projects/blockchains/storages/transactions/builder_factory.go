package transactions

// BuilderFactory represents a transactions builder factory
type BuilderFactory interface {
	Create() Builder
}
