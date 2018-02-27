package infrastructure

import (
	"encoding/json"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

// JsDataForTests represents a structure for tests
type JsDataForTests struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests(t *testing.T) *Transaction {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	trs := createTransaction(&id, js, createdOn)
	return trs.(*Transaction)
}
