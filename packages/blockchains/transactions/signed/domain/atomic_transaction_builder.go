package domain

import (
	"time"

	met "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	uuid "github.com/satori/go.uuid"
)

// AtomicTransactionBuilder represents an AtomicTransaction builder
type AtomicTransactionBuilder interface {
	Create() AtomicTransactionBuilder
	WithID(id *uuid.UUID) AtomicTransactionBuilder
	WithMetaData(meta met.MetaData) AtomicTransactionBuilder
	WithTransactions(trs transactions.Transactions) AtomicTransactionBuilder
	WithSignature(sig users.Signature) AtomicTransactionBuilder
	CreatedOn(createdOn time.Time) AtomicTransactionBuilder
	Now() (AtomicTransaction, error)
}
