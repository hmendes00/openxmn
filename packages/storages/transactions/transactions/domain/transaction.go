package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// Transaction represents a stored signed transaction
type Transaction interface {
	GetHashTree() hashtrees.HashTree
	GetSignature() users.Signature
	GetTrs() stored_chunks.Chunks
	CreatedOn() time.Time
}
