package domain

import (
	"time"

	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// TransactionBuilder represents a stored signed transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithSignature(sig users.Signature) TransactionBuilder
	WithTrs(chks stored_chunks.Chunks) TransactionBuilder
	CreatedOn(ts time.Time) TransactionBuilder
	Now() (Transaction, error)
}
