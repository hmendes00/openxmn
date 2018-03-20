package aggregated

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
	aggregated "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	users "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/users"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/blockchains/metadata"
	concrete_users "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

type signedTransactionsBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	trs                    aggregated.Transactions
	sig                    users.Signature
	createdOn              *time.Time
}

func createSignedTransactionsBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) aggregated.SignedTransactionsBuilder {
	out := signedTransactionsBuilder{
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

// Create initializes the builder
func (build *signedTransactionsBuilder) Create() aggregated.SignedTransactionsBuilder {
	build.id = nil
	build.met = nil
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

// WithMetaData adds the MetaData to the builder
func (build *signedTransactionsBuilder) WithMetaData(met metadata.MetaData) aggregated.SignedTransactionsBuilder {
	build.met = met
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

	if build.trs == nil {
		return nil, errors.New("the Transactions is mandatory in order to build a SignedTransactions")
	}

	if build.sig == nil {
		return nil, errors.New("the user Signature is mandatory in order to build a SignedTransactions")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a SignedTransactions")
		}

		if build.createdOn == nil {
			return nil, errors.New("the createdOn time is mandatory in order to build a SignedTransactions")
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

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a SignedTransactions instance")
	}

	out := createSignedTransactions(build.met.(*concrete_metadata.MetaData), build.trs.(*Transactions), build.sig.(*concrete_users.Signature))
	return out, nil

}
