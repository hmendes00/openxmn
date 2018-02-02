package infrastructure

import (
	signed_transactions "github.com/XMNBlockchain/core/packages/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/transactions/transactions/domain"
	concrete_transactions "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// Transaction represents the concrete signed transaction
type Transaction struct {
	ID  *uuid.UUID                         `json:"id"`
	Trs *concrete_transactions.Transaction `json:"transaction"`
	Sig *concrete_users.Signature          `json:"signature"`
}

func createTransaction(id *uuid.UUID, trs *concrete_transactions.Transaction, sig *concrete_users.Signature) signed_transactions.Transaction {
	out := Transaction{
		ID:  id,
		Trs: trs,
		Sig: sig,
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
