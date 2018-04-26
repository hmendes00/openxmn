package write

import (
	"errors"
	"fmt"

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
func (db *Project) Insert(proj *objects.Project) {
	db.projects[proj.Met.GetID().String()] = proj
}

// Update updates an existing project
func (db *Project) Update(original *objects.Project, new *objects.Project) error {
	delErr := db.Delete(original)
	if delErr != nil {
		return delErr
	}

	db.Insert(new)
	return nil
}

// Delete deletes an existing project
func (db *Project) Delete(proj *objects.Project) error {
	idAsString := proj.Met.GetID().String()
	if _, ok := db.projects[idAsString]; ok {
		delete(db.projects, idAsString)
		return nil
	}

	str := fmt.Sprintf("the project (ID: %s) could not be found", idAsString)
	return errors.New(str)
}
