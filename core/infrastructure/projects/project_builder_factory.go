package projects

import (
	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
)

// ProjectBuilderFactory represents a concrete ProjectBuilderFactory implementation
type ProjectBuilderFactory struct {
}

// CreateProjectBuilderFactory creates a new ProjectBuilderFactory instance
func CreateProjectBuilderFactory() projects.ProjectBuilderFactory {
	out := ProjectBuilderFactory{}
	return &out
}

// Create creates a ProjectBuilder instance
func (fac *ProjectBuilderFactory) Create() projects.ProjectBuilder {
	out := createProjectBuilder()
	return out
}
