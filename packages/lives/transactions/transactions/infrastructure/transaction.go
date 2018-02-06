package infrastructure

import (
	"time"

	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body"
	concrete_body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents the concrete transaction
type Transaction struct {
	ID    *uuid.UUID          `json:"id"`
	Karma int                 `json:"karma"`
	Bod   *concrete_body.Body `json:"body"`
	CrOn  time.Time           `json:"created_on"`
}

func createTransaction(id *uuid.UUID, karma int, body *concrete_body.Body, createdOn time.Time) trs.Transaction {
	out := Transaction{
		ID:    id,
		Bod:   body,
		Karma: karma,
		CrOn:  createdOn,
	}

	return &out
}

// GetID returns the ID of the transaction
func (trs *Transaction) GetID() *uuid.UUID {
	return trs.ID
}

// GetBody returns the Body of the transaction
func (trs *Transaction) GetBody() body.Body {
	return trs.Bod
}

// GetKarma returns the karma of the transaction
func (trs *Transaction) GetKarma() int {
	return trs.Karma
}

// CreatedOn returns the cresation time of the transaction
func (trs *Transaction) CreatedOn() time.Time {
	return trs.CrOn
}
