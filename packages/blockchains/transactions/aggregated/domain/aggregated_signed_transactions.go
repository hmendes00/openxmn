package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
)

// AggregatedSignedTransactions represents multiple SignedTransactions aggregated together and ordered by an HashMap
type AggregatedSignedTransactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() []SignedTransactions
}
