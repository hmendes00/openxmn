package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	uuid "github.com/satori/go.uuid"
)

// SignedTransactionsBuilder represents the SignedTransactions builder
type SignedTransactionsBuilder interface {
	Create() SignedTransactionsBuilder
	WithID(id *uuid.UUID) SignedTransactionsBuilder
	WithMetaData(met metadata.MetaData) SignedTransactionsBuilder
	WithTransactions(trs Transactions) SignedTransactionsBuilder
	WithSignature(sig users.Signature) SignedTransactionsBuilder
	CreatedOn(ts time.Time) SignedTransactionsBuilder
	Now() (SignedTransactions, error)
}
