package wallets

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Wallet represents a user wallet
type Wallet interface {
	GetMetaData() metadata.MetaData
	GetOwner() users.User
	GetToken() tokens.Token
	GetAmount() float64
}
