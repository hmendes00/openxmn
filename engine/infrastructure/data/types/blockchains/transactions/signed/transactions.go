package signed

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
)

// Transactions represents []Transaction ordered by an HashMap
type Transactions struct {
	Met *concrete_metadata.MetaData `json:"metadata"`
	Trs []*Transaction              `json:"transactions"`
}

func createTransactions(met *concrete_metadata.MetaData, trs []*Transaction) signed_transactions.Transactions {
	out := Transactions{
		Met: met,
		Trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *Transactions) GetMetaData() metadata.MetaData {
	return trs.Met
}

// GetTransactions returns the []Transaction
func (trs *Transactions) GetTransactions() []signed_transactions.Transaction {
	out := []signed_transactions.Transaction{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}

	return out
}
