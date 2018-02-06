package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithHashTree(ht hashtrees.HashTree) BlockBuilder
	WithSignature(sig users.Signature) BlockBuilder
	WithTransactions(trs stored_aggregated_transactions.Transactions) BlockBuilder
	WithNeededKarma(neededKarma int) BlockBuilder
	CreatedOn(ts time.Time) BlockBuilder
	Now() (Block, error)
}
