package signed

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
	met "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
	trs "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions/signed"
	users "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/users"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/blockchains/metadata"
	concrete_transactions "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/blockchains/transactions"
	concrete_users "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

type atomicTransactionBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory met.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    met.MetaData
	trs                    trs.Transactions
	sig                    users.Signature
	createdOn              *time.Time
}

func createAtomicTransactionBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory met.MetaDataBuilderFactory) signed_transactions.AtomicTransactionBuilder {
	out := atomicTransactionBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:        nil,
		met:       nil,
		trs:       nil,
		sig:       nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the AtomicTransactionBuilder instance
func (build *atomicTransactionBuilder) Create() signed_transactions.AtomicTransactionBuilder {
	build.id = nil
	build.met = nil
	build.trs = nil
	build.sig = nil
	return build
}

// WithID adds an ID to the AtomicTransactionBuilder
func (build *atomicTransactionBuilder) WithID(id *uuid.UUID) signed_transactions.AtomicTransactionBuilder {
	build.id = id
	return build
}

// WithMetaData adds a MetaData to the AtomicTransactionBuilder
func (build *atomicTransactionBuilder) WithMetaData(meta met.MetaData) signed_transactions.AtomicTransactionBuilder {
	build.met = meta
	return build
}

// WithTransactions adds transactions to the AtomicTransactionBuilder
func (build *atomicTransactionBuilder) WithTransactions(trs trs.Transactions) signed_transactions.AtomicTransactionBuilder {
	build.trs = trs
	return build
}

// WithSignature adds a user signature to the AtomicTransactionBuilder
func (build *atomicTransactionBuilder) WithSignature(sig users.Signature) signed_transactions.AtomicTransactionBuilder {
	build.sig = sig
	return build
}

// CreatedOn adds a creation time to the AtomicTransactionBuilder
func (build *atomicTransactionBuilder) CreatedOn(ts time.Time) signed_transactions.AtomicTransactionBuilder {
	build.createdOn = &ts
	return build
}

// Now builds a new AtomicTransaction instance
func (build *atomicTransactionBuilder) Now() (signed_transactions.AtomicTransaction, error) {

	if build.trs == nil {
		return nil, errors.New("the []transaction are mandatory in order to build an AtomicTransaction instance")
	}

	if build.sig == nil {
		return nil, errors.New("the user signature is mandatory in order to build an AtomicTransaction instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build an AtomicTransaction instance")
		}

		if build.createdOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build an AtomicTransaction instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.createdOn.UnixNano()))),
			build.trs.GetMetaData().GetHashTree().GetHash().Get(),
			build.sig.GetMetaData().GetHashTree().GetHash().Get(),
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.createdOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	out := createAtomicTransaction(build.met.(*concrete_metadata.MetaData), build.trs.(*concrete_transactions.Transactions), build.sig.(*concrete_users.Signature))
	return out, nil
}
