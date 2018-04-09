package wallets

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// WalletBuilder represents a wallet builder
type WalletBuilder interface {
	Create() WalletBuilder
	WithMetaData(met metadata.MetaData) WalletBuilder
	WithOwner(owner users.User) WalletBuilder
	WithToken(tok tokens.Token) WalletBuilder
	WithAmount(amount float64) WalletBuilder
	Now() (Wallet, error)
}
