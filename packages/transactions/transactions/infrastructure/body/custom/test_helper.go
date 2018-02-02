package custom

import (
	"encoding/json"
	"testing"

	uuid "github.com/satori/go.uuid"
)

// InstanceForTests represents a custom instance for tests
type InstanceForTests struct {
	ID          *uuid.UUID `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}

// CreateCreateForTests creates a Create for tests
func CreateCreateForTests(t *testing.T) *Create {
	//variables:
	id := uuid.NewV4()
	insID := uuid.NewV4()
	ins := InstanceForTests{
		ID:          &insID,
		Name:        "some name",
		Description: "this is some custom description",
	}

	js, _ := json.Marshal(ins)
	cr := createCreate(&id, js)
	return cr.(*Create)
}

// CreateCustomWithCreateForTests creates a Custom for tests, using a create transaction
func CreateCustomWithCreateForTests(t *testing.T) *Custom {
	//variables:
	cr := CreateCreateForTests(t)

	cu := createCustomWithCreate(cr)
	return cu.(*Custom)
}
