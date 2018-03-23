package domain

import (
	"time"

	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
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
