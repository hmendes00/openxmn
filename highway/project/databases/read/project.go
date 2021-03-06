package read

import (
	"errors"
	"fmt"

	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	uuid "github.com/satori/go.uuid"
)

// Project represents a project read database
type Project struct {
	projs map[string]*objects.Project
}

// CreateProject creates a new Project instance
func CreateProject(projs map[string]*objects.Project) *Project {
	out := Project{
		projs: projs,
	}

	return &out
}

// RetrieveByID retrieves a project by its ID
func (db *Project) RetrieveByID(id *uuid.UUID) (*objects.Project, error) {
	idAsString := id.String()
	if oneProj, ok := db.projs[idAsString]; ok {
		return oneProj, nil
	}

	str := fmt.Sprintf("the project (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

// CanUpdate verifies if a given user can update the given project
func (db *Project) CanUpdate(proj *objects.Project, user users.User) bool {
	return true
}

// CanDelete verifies if a given user can delete the given project
func (db *Project) CanDelete(proj *objects.Project, user users.User) bool {
	return true
}
