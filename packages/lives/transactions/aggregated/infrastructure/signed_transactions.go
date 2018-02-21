package infrastructure

import (
	"time"

	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// SignedTransactions represents the concrete signed transactions
type SignedTransactions struct {
	ID   *uuid.UUID                `json:"id"`
	Trs  *Transactions             `json:"transactions"`
	Sig  *concrete_users.Signature `json:"signature"`
	CrOn time.Time                 `json:"created_on"`
}

func createSignedTransactions(id *uuid.UUID, trs *Transactions, sig *concrete_users.Signature, createdOn time.Time) aggregated.SignedTransactions {
	out := SignedTransactions{
		ID:   id,
		Trs:  trs,
		Sig:  sig,
		CrOn: createdOn,
	}

	return &out
}

// GetID returns the ID of the signed transactions
func (trs *SignedTransactions) GetID() *uuid.UUID {
	return trs.ID
}

// GetTrs returns the Transactions instance
func (trs *SignedTransactions) GetTrs() aggregated.Transactions {
	return trs.Trs
}

// GetSignature returns the user Signature
func (trs *SignedTransactions) GetSignature() users.Signature {
	return trs.Sig
}

// CreatedOn returns the creation date time
func (trs *SignedTransactions) CreatedOn() time.Time {
	return trs.CrOn
}
