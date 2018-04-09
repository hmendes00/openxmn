package bills

import (
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

// PayerBuilder represents a payer builder
type PayerBuilder interface {
	Create() PayerBuilder
	WithOutgoingBandwidthInBytes(outBw int) PayerBuilder
	WithWallet(wal wallets.Wallet) PayerBuilder
	Now() (Payer, error)
}
