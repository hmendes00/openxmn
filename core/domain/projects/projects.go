package projects

import uuid "github.com/satori/go.uuid"

// Projects represents a list of projects
type Projects interface {
	IsEmpty() bool
	GetAmount() int
	GetByID(id *uuid.UUID) Project
	GetProjects(index int, amount int) []Project
	GetAllProjects() []Project
}
