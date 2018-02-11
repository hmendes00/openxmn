package infrastructure

import (
	"errors"
	"time"

	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type transactionBuilder struct {
	id   *uuid.UUID
	trs  trs.Transaction
	sig  users.Signature
	crOn *time.Time
}

func createTransactionBuilder() signed_transactions.TransactionBuilder {
	out := transactionBuilder{
		id:   nil,
		trs:  nil,
		sig:  nil,
		crOn: nil,
	}

	return &out
}

// Create initializes the TransactionBuilder instance
func (build *transactionBuilder) Create() signed_transactions.TransactionBuilder {
	build.id = nil
	build.trs = nil
	build.sig = nil
	build.crOn = nil
	return build
}

func (build *transactionBuilder) WithID(id *uuid.UUID) signed_transactions.TransactionBuilder {
	build.id = id
	return build
}

// WithTransaction adds a Transaction to the signed TransactionBuilder
func (build *transactionBuilder) WithTransaction(trs trs.Transaction) signed_transactions.TransactionBuilder {
	build.trs = trs
	return build
}

// WithSignature adds a user Signature to the signed TransactionBuilder
func (build *transactionBuilder) WithSignature(sig users.Signature) signed_transactions.TransactionBuilder {
	build.sig = sig
	return build
}

// CreatedOn adds the creation time to the TransactionBuilder
func (build *transactionBuilder) CreatedOn(ts time.Time) signed_transactions.TransactionBuilder {
	build.crOn = &ts
	return build
}

// Now builds a signed Transaction instance
func (build *transactionBuilder) Now() (signed_transactions.Transaction, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a signed Transaction instance")
	}

	if build.trs == nil {
		return nil, errors.New("the transaction is mandatory in order to build a signed Transaction instance")
	}

	if build.sig == nil {
		return nil, errors.New("the user signature is mandatory in order to build a signed Transaction instance")
	}

	if build.crOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a signed Transaction instance")
	}

	out := createTransaction(build.id, build.trs.(*concrete_transactions.Transaction), build.sig.(*concrete_users.Signature), *build.crOn)
	return out, nil
}
