package domain

import (
	"time"

	met "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

// TransactionBuilder represents a signed transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithID(id *uuid.UUID) TransactionBuilder
	WithMetaData(meta met.MetaData) TransactionBuilder
	WithTransaction(trs transactions.Transaction) TransactionBuilder
	WithSignature(sig users.Signature) TransactionBuilder
	CreatedOn(ts time.Time) TransactionBuilder
	Now() (Transaction, error)
}
