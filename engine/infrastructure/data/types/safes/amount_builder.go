package safes

import (
	"errors"

	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	uuid "github.com/satori/go.uuid"
)

type amountBuilder struct {
	tokID  *uuid.UUID
	amount float64
}

func createAmountBuilder() safes.AmountBuilder {
	out := amountBuilder{
		tokID:  nil,
		amount: float64(0),
	}

	return &out
}

// Create initializes the AmountBuilder
func (build *amountBuilder) Create() safes.AmountBuilder {
	build.tokID = nil
	build.amount = float64(0)
	return build
}

// WithTokenID adds a tokenID to the AmountBuilder
func (build *amountBuilder) WithTokenID(id *uuid.UUID) safes.AmountBuilder {
	build.tokID = id
	return build
}

// WithAmount adds an amount to the AmountBuilder
func (build *amountBuilder) WithAmount(amount float64) safes.AmountBuilder {
	build.amount = amount
	return build
}

// Now builds a new Amount instance
func (build *amountBuilder) Now() (safes.Amount, error) {
	if build.tokID == nil {
		return nil, errors.New("the ID is mandatory in order to build an amount")
	}

	if build.amount <= 0 {
		return nil, errors.New("the amount must be greater than 0 in order to build an amount")
	}

	out := createAmount(build.tokID, build.amount)
	return out, nil
}
