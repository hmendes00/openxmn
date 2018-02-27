package infrastructure

import (
	"errors"
	"time"

	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	uuid "github.com/satori/go.uuid"
)

type transactionBuilder struct {
	id        *uuid.UUID
	js        []byte
	createdOn *time.Time
}

func createTransactionBuilder() trs.TransactionBuilder {
	out := transactionBuilder{}
	return &out
}

// Create initializes the transactionBuilder
func (build *transactionBuilder) Create() trs.TransactionBuilder {
	build.id = nil
	build.js = nil
	build.createdOn = nil
	return build
}

// WithID adds an ID to the transactionBuilder
func (build *transactionBuilder) WithID(id *uuid.UUID) trs.TransactionBuilder {
	build.id = id
	return build
}

// WithJSON adds JSON data to the transactionBuilder
func (build *transactionBuilder) WithJSON(js []byte) trs.TransactionBuilder {
	build.js = js
	return build
}

// CreatedOn adds the creation time to the transactionBuilder
func (build *transactionBuilder) CreatedOn(time time.Time) trs.TransactionBuilder {
	build.createdOn = &time
	return build
}

// Now build a new transaction instance
func (build *transactionBuilder) Now() (trs.Transaction, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a transaction instance")
	}

	if build.js == nil {
		return nil, errors.New("the json data is mandatory in order to build a transaction instance")
	}

	if build.createdOn == nil {
		return nil, errors.New("the createdOn is mandatory in order to build a transaction instance")
	}

	out := createTransaction(build.id, build.js, *build.createdOn)
	return out, nil
}
