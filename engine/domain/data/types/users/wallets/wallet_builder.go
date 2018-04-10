package wallets

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// WalletBuilder represents a wallet builder
type WalletBuilder interface {
	Create() WalletBuilder
	WithID(id *uuid.UUID) WalletBuilder
	WithMetaData(met metadata.MetaData) WalletBuilder
	WithOwner(owner users.User) WalletBuilder
	WithToken(tok tokens.Token) WalletBuilder
	WithAmount(amount float64) WalletBuilder
	CreatedOn(crOn time.Time) WalletBuilder
	LastUpdatedOn(crOn time.Time) WalletBuilder
	Now() (Wallet, error)
}
