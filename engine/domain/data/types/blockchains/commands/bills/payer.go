package bills

import (
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

// Payer represents an entity that pays a bill
type Payer interface {
	GetOutgoingBandwidthInBytes() int
	GetWallet() wallets.Wallet
}
