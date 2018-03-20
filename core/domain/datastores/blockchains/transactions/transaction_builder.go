package domain

import (
	"time"

	met "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

// TransactionBuilder represents the Transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithID(id *uuid.UUID) TransactionBuilder
	WithMetaData(meta met.MetaData) TransactionBuilder
	WithJSON(data []byte) TransactionBuilder
	CreatedOn(time time.Time) TransactionBuilder
	Now() (Transaction, error)
}
