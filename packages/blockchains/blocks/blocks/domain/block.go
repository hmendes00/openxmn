package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
)

// Block represents multiple SignedTransactions aggregated together and ordered by an HashMap
type Block interface {
	GetMetaData() metadata.MetaData
	GetTransactions() []aggregated.SignedTransactions
}
