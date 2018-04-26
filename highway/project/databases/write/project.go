package write

import (
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Project represents a project write database
type Project struct {
	projects map[string]*objects.Project
}

// CreateProject creates a new Project instance
func CreateProject(projects map[string]*objects.Project) *Project {
	out := Project{
		projects: projects,
	}

	return &out
}

// Insert inserts a new project
func (db *Project) Insert(proj *objects.Project) error {
	return nil
}

// Update updates an existing project
func (db *Project) Update(original *objects.Project, new *objects.Project) error {
	return nil
}

// Delete deletes an existing project
func (db *Project) Delete(proj *objects.Project) error {
	return nil
}
