package domain

import (
	"time"

	met "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
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