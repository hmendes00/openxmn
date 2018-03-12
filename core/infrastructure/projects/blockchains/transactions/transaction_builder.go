package transactions

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
	met "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	trs "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions"
	concrete_met "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

type transactionBuilder struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory met.MetaDataBuilderFactory
	meta                   met.MetaData
	id                     *uuid.UUID
	js                     []byte
	createdOn              *time.Time
}

func createTransactionBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory met.MetaDataBuilderFactory) trs.TransactionBuilder {
	out := transactionBuilder{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		meta:      nil,
		id:        nil,
		js:        nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the transactionBuilder
func (build *transactionBuilder) Create() trs.TransactionBuilder {
	build.id = nil
	build.meta = nil
	build.js = nil
	build.createdOn = nil
	return build
}

// WithMetaData adds a metadata instance to the transactionBuilder
func (build *transactionBuilder) WithMetaData(meta met.MetaData) trs.TransactionBuilder {
	build.meta = meta
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

	if build.js == nil {
		return nil, errors.New("the json data is mandatory in order to build a transaction instance")
	}

	if build.meta == nil {

		if build.id == nil {
			return nil, errors.New("since there is no metadata, the ID is mandatory in order to build a transaction instance")
		}

		if build.createdOn == nil {
			return nil, errors.New("since there is no metadata, the createdOn is mandatory in order to build a transaction instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.createdOn.UnixNano()))),
			build.js,
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		meta, metaErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.createdOn).Now()
		if metaErr != nil {
			return nil, metaErr
		}

		build.meta = meta
	}

	if build.meta == nil {
		return nil, errors.New("the metadata is mandatory in order to build a transaction instance")
	}

	out := createTransaction(build.meta.(*concrete_met.MetaData), build.js)
	return out, nil
}
