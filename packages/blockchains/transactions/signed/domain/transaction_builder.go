package domain

import (
	"time"

	met "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
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
