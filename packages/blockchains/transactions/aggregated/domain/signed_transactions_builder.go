package domain

import (
	"time"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	uuid "github.com/satori/go.uuid"
)

// SignedTransactionsBuilder represents the SignedTransactions builder
type SignedTransactionsBuilder interface {
	Create() SignedTransactionsBuilder
	WithID(id *uuid.UUID) SignedTransactionsBuilder
	WithUserID(usrID *uuid.UUID) SignedTransactionsBuilder
	WithPrivateKey(pk cryptography.PrivateKey) SignedTransactionsBuilder
	WithTransactions(trs Transactions) SignedTransactionsBuilder
	WithSignature(sig users.Signature) SignedTransactionsBuilder
	CreatedOn(ts time.Time) SignedTransactionsBuilder
	Now() (SignedTransactions, error)
}
