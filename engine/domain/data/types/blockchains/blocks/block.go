package domain

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	aggregated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed/aggregated"
)

// Block represents multiple SignedTransactions aggregated together and ordered by an HashMap
type Block interface {
	GetMetaData() metadata.MetaData
	GetTransactions() []aggregated.SignedTransactions
}
