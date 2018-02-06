package infrastructure

import (
	"errors"
	"time"

	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body"
	concrete_body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body"
	uuid "github.com/satori/go.uuid"
)

type transactionBuilder struct {
	id        *uuid.UUID
	bod       body.Body
	karma     int
	createdOn *time.Time
}

func createTransactionBuilder() trs.TransactionBuilder {
	out := transactionBuilder{}
	return &out
}

// Create initializes the transactionBuilder
func (build *transactionBuilder) Create() trs.TransactionBuilder {
	build.id = nil
	build.bod = nil
	build.karma = -1
	build.createdOn = nil
	return build
}

// WithID adds an ID to the transactionBuilder
func (build *transactionBuilder) WithID(id *uuid.UUID) trs.TransactionBuilder {
	build.id = id
	return build
}

// WithBody adds a Body to the transactionBuilder
func (build *transactionBuilder) WithBody(bod body.Body) trs.TransactionBuilder {
	build.bod = bod
	return build
}

// WithKarma adds karma to the transactionBuilder
func (build *transactionBuilder) WithKarma(karma int) trs.TransactionBuilder {
	build.karma = karma
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

	if build.bod == nil {
		return nil, errors.New("the body is mandatory in order to build a transaction instance")
	}

	if build.karma <= 0 {
		return nil, errors.New("the karma must be a positive value (greater than 0)")
	}

	if build.createdOn == nil {
		return nil, errors.New("the createdOn is mandatory in order to build a transaction instance")
	}

	out := createTransaction(build.id, build.karma, build.bod.(*concrete_body.Body), *build.createdOn)
	return out, nil
}
