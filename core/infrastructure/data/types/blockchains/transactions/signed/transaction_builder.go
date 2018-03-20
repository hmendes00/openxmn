package signed

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/hashtrees"
	met "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	trs "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
	concrete_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/transactions"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

type transactionBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory met.MetaDataBuilderFactory
	id                     *uuid.UUID
	meta                   met.MetaData
	trs                    trs.Transaction
	sig                    users.Signature
	crOn                   *time.Time
}

func createTransactionBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory met.MetaDataBuilderFactory) signed_transactions.TransactionBuilder {
	out := transactionBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:   nil,
		meta: nil,
		trs:  nil,
		sig:  nil,
		crOn: nil,
	}

	return &out
}

// Create initializes the TransactionBuilder instance
func (build *transactionBuilder) Create() signed_transactions.TransactionBuilder {
	build.id = nil
	build.meta = nil
	build.trs = nil
	build.sig = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the TransactionBuilder instance
func (build *transactionBuilder) WithID(id *uuid.UUID) signed_transactions.TransactionBuilder {
	build.id = id
	return build
}

// WithMetaData adds a MetaData to the TransactionBuilder instance
func (build *transactionBuilder) WithMetaData(meta met.MetaData) signed_transactions.TransactionBuilder {
	build.meta = meta
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

	if build.trs == nil {
		return nil, errors.New("the transaction is mandatory in order to build a signed Transaction instance")
	}

	if build.sig == nil {
		return nil, errors.New("the user signature is mandatory in order to build a signed Transaction instance")
	}

	if build.meta == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a signed Transaction instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a signed Transaction instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.crOn.UnixNano()))),
			build.trs.GetMetaData().GetHashTree().GetHash().Get(),
			build.sig.GetMetaData().GetHashTree().GetHash().Get(),
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.crOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.meta = met
	}

	out := createTransaction(build.meta.(*concrete_metadata.MetaData), build.trs.(*concrete_transactions.Transaction), build.sig.(*concrete_users.Signature))
	return out, nil
}
