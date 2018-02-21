package infrastructure

import (
	"errors"
	"time"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type signedTransactionsBuilder struct {
	sigBuilderFactory users.SignatureBuilderFactory
	id                *uuid.UUID
	usrID             *uuid.UUID
	pk                cryptography.PrivateKey
	trs               aggregated.Transactions
	sig               users.Signature
	createdOn         *time.Time
}

func createSignedTransactionsBuilder(sigBuilderFactory users.SignatureBuilderFactory) aggregated.SignedTransactionsBuilder {
	out := signedTransactionsBuilder{
		sigBuilderFactory: sigBuilderFactory,
		id:                nil,
		usrID:             nil,
		pk:                nil,
		trs:               nil,
		sig:               nil,
		createdOn:         nil,
	}

	return &out
}

// Create initializes the builder
func (build *signedTransactionsBuilder) Create() aggregated.SignedTransactionsBuilder {
	build.id = nil
	build.usrID = nil
	build.pk = nil
	build.trs = nil
	build.sig = nil
	build.createdOn = nil
	return build
}

// WithID adds an ID instance to the builder
func (build *signedTransactionsBuilder) WithID(id *uuid.UUID) aggregated.SignedTransactionsBuilder {
	build.id = id
	return build
}

// WithUserID adds a userID to the builder
func (build *signedTransactionsBuilder) WithUserID(usrID *uuid.UUID) aggregated.SignedTransactionsBuilder {
	build.usrID = usrID
	return build
}

// WithPrivateKey adds a PrivateKey instance to the builder
func (build *signedTransactionsBuilder) WithPrivateKey(pk cryptography.PrivateKey) aggregated.SignedTransactionsBuilder {
	build.pk = pk
	return build
}

// WithTransactions adds a Transactions instance to the builder
func (build *signedTransactionsBuilder) WithTransactions(trs aggregated.Transactions) aggregated.SignedTransactionsBuilder {
	build.trs = trs
	return build
}

// WithSignature adds a user Signature instance to the builder
func (build *signedTransactionsBuilder) WithSignature(sig users.Signature) aggregated.SignedTransactionsBuilder {
	build.sig = sig
	return build
}

// WithSignature adds a createdOn time to the builder
func (build *signedTransactionsBuilder) CreatedOn(ts time.Time) aggregated.SignedTransactionsBuilder {
	build.createdOn = &ts
	return build
}

// Now builds a new SignedTransactions instance
func (build *signedTransactionsBuilder) Now() (aggregated.SignedTransactions, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a SignedTransactions")
	}

	if build.createdOn == nil {
		return nil, errors.New("the createdOn time is mandatory in order to build a SignedTransactions")
	}

	if build.trs == nil {
		return nil, errors.New("the Transactions is mandatory in order to build a SignedTransactions")
	}

	if build.sig == nil && build.usrID == nil {
		return nil, errors.New("the user Signature or the userID is mandatory in order to build a SignedTransactions")
	}

	if build.sig == nil && build.usrID != nil && build.pk != nil {
		sig, sigErr := build.sigBuilderFactory.Create().Create().WithUserID(build.usrID).WithInterface(build.trs).WithPrivateKey(build.pk).Now()
		if sigErr != nil {
			return nil, sigErr
		}

		build.sig = sig
	}

	out := createSignedTransactions(build.id, build.trs.(*Transactions), build.sig.(*concrete_users.Signature), *build.createdOn)
	return out, nil

}
