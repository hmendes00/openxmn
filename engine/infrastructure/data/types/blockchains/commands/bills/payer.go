package bills

import (
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

type payer struct {
	outBwInBytes int
	wal          wallets.Wallet
}

func createPayer(outBwInBytes int, wal wallets.Wallet) bills.Payer {
	out := payer{
		outBwInBytes: outBwInBytes,
		wal:          wal,
	}

	return &out
}

// GetOutgoingBandwidthInBytes returns the outgoing bandwidth in bytes
func (pay *payer) GetOutgoingBandwidthInBytes() int {
	return pay.outBwInBytes
}

// GetWallet returns the wallet
func (pay *payer) GetWallet() wallets.Wallet {
	return pay.wal
}
