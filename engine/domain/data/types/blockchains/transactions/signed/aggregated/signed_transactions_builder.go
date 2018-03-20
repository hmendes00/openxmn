package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/users"
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
