package projects

// ProjectBuilderFactory represents a project builder factory
type ProjectBuilderFactory interface {
	Create() ProjectBuilder
}
