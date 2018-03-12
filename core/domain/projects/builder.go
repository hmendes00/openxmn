package projects

// Builder represents a projects builder
type Builder interface {
	Create() Builder
	WithProjects(projs []Project) Builder
	Now() Projects
}
