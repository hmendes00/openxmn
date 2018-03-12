package projects

import (
	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
	uuid "github.com/satori/go.uuid"
)

// Projects represents a concrete projects implementation
type Projects struct {
	List []*Project `json:"projects"`
	Mp   map[string]*Project
}

func createProjects(lst []*Project) projects.Projects {

	mp := map[string]*Project{}
	for _, oneProj := range lst {
		idAsString := oneProj.GetID().String()
		mp[idAsString] = oneProj
	}

	out := Projects{
		List: lst,
		Mp:   mp,
	}

	return &out
}

// IsEmpty returns true if the list is empty, false otherwise
func (projs *Projects) IsEmpty() bool {
	return len(projs.List) <= 0
}

// GetAmount returns the amount of projects
func (projs *Projects) GetAmount() int {
	return len(projs.Mp)
}

// GetByID returns a Project associated to an ID
func (projs *Projects) GetByID(id *uuid.UUID) projects.Project {
	idAsString := id.String()
	if oneProj, ok := projs.Mp[idAsString]; ok {
		return oneProj
	}

	return nil
}

// GetProjects returns the []Project indexed inside the index and amount indexes
func (projs *Projects) GetProjects(index int, amount int) []projects.Project {

	//max to index:
	toIndex := index + amount
	totalAmount := projs.GetAmount()
	if toIndex > totalAmount {
		toIndex = totalAmount
	}

	//slice then return:
	sliced := projs.List[index:toIndex]
	out := []projects.Project{}
	for _, oneProj := range sliced {
		out = append(out, oneProj)
	}

	return out
}

// GetAllProjects returns all the projects
func (projs *Projects) GetAllProjects() []projects.Project {
	out := []projects.Project{}
	for _, oneProj := range projs.List {
		out = append(out, oneProj)
	}

	return out
}
