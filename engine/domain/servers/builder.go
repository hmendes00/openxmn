package projects

// Builder represents a servers builder
type Builder interface {
	Create() Builder
	WithServers(servs []Server) Builder
	Now() Servers
}
