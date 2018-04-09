package bills

import (
	"errors"

	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
	user_wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

type payerBuilder struct {
	outBwInBytes int
	wal          user_wallets.Wallet
}

func createPayerBuilder() bills.PayerBuilder {
	out := payerBuilder{
		outBwInBytes: 0,
		wal:          nil,
	}

	return &out
}

// Create initializes a PayerBuilder instance
func (build *payerBuilder) Create() bills.PayerBuilder {
	build.outBwInBytes = 0
	build.wal = nil
	return build
}

// WithOutgoingBandwidthInBytes adds outgoing bandwidth in bytes to the PayerBuilder instance
func (build *payerBuilder) WithOutgoingBandwidthInBytes(outBw int) bills.PayerBuilder {
	build.outBwInBytes = outBw
	return build
}

// WithWallet adds a wallet to the PayerBuilder instance
func (build *payerBuilder) WithWallet(wal wallets.Wallet) bills.PayerBuilder {
	build.wal = wal
	return build
}

// Now builds a new Payer instance
func (build *payerBuilder) Now() (bills.Payer, error) {
	if build.outBwInBytes == 0 {
		return nil, errors.New("the outgoing bandwidth in bytes is mandatory in order to build a Payer instance")
	}

	if build.wal == nil {
		return nil, errors.New("the wallet is mandatory in order to build a Payer instance")
	}

	out := createPayer(build.outBwInBytes, build.wal)
	return out, nil
}
