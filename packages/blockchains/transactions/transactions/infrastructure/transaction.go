package infrastructure

import (
	"time"

	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents the concrete transaction
type Transaction struct {
	ID   *uuid.UUID `json:"id"`
	JS   []byte     `json:"json"`
	CrOn time.Time  `json:"created_on"`
}

func createTransaction(id *uuid.UUID, js []byte, createdOn time.Time) trs.Transaction {
	out := Transaction{
		ID:   id,
		JS:   js,
		CrOn: createdOn,
	}

	return &out
}

// GetID returns the ID of the transaction
func (trs *Transaction) GetID() *uuid.UUID {
	return trs.ID
}

// GetJSON returns the json data of the transaction
func (trs *Transaction) GetJSON() []byte {
	return trs.JS
}

// CreatedOn returns the cresation time of the transaction
func (trs *Transaction) CreatedOn() time.Time {
	return trs.CrOn
}
