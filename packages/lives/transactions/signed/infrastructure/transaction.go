package infrastructure

import (
	"time"

	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents the concrete signed transaction
type Transaction struct {
	ID   *uuid.UUID                         `json:"id"`
	Trs  *concrete_transactions.Transaction `json:"transaction"`
	Sig  *concrete_users.Signature          `json:"signature"`
	CrOn time.Time                          `json:"created_on"`
}

func createTransaction(id *uuid.UUID, trs *concrete_transactions.Transaction, sig *concrete_users.Signature, createdOn time.Time) signed_transactions.Transaction {
	out := Transaction{
		ID:   id,
		Trs:  trs,
		Sig:  sig,
		CrOn: createdOn,
	}

	return &out
}

// GetID returns the ID
func (trs *Transaction) GetID() *uuid.UUID {
	return trs.ID
}

// GetTrs returns the Transaction
func (trs *Transaction) GetTrs() trs.Transaction {
	return trs.Trs
}

// GetSignature returns the user signature
func (trs *Transaction) GetSignature() users.Signature {
	return trs.Sig
}

// CreatedOn returns the creation time
func (trs *Transaction) CreatedOn() time.Time {
	return trs.CrOn
}
