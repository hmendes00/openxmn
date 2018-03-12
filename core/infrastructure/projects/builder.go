package projects

import (
	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
)

type builder struct {
	list []projects.Project
}

func createBuilder() projects.Builder {
	out := builder{
		list: []projects.Project{},
	}

	return &out
}

// Create initializes the builder
func (build *builder) Create() projects.Builder {
	build.list = []projects.Project{}
	return build
}

// WithProjects adds []Project to the builder
func (build *builder) WithProjects(projs []projects.Project) projects.Builder {
	build.list = projs
	return build
}

// Now builds a new Projects instance
func (build *builder) Now() projects.Projects {
	in := []*Project{}
	for _, oneProj := range build.list {
		in = append(in, oneProj.(*Project))
	}

	out := createProjects(in)
	return out
}
