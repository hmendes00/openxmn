package domain

import (
	"time"

	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// AtomicTransactionBuilder represents a stored atomic signed TransactionBuilder
type AtomicTransactionBuilder interface {
	Create() AtomicTransactionBuilder
	WithSignature(isg users.Signature) AtomicTransactionBuilder
	WithTrs(trs []stored_chunks.Chunks) AtomicTransactionBuilder
	CreatedOn(ts time.Time) AtomicTransactionBuilder
	Now() (AtomicTransaction, error)
}
