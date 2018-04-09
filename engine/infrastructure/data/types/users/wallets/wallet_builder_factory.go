package wallets

import (
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

// WalletBuilderFactory represents a concrete WalletBuilderFactory implementation
type WalletBuilderFactory struct {
}

// CreateWalletBuilderFactory creates a new WalletBuilderFactory implementation
func CreateWalletBuilderFactory() wallets.WalletBuilderFactory {
	out := WalletBuilderFactory{}
	return &out
}

// Create creates a new WalletBuilder instance
func (fac *WalletBuilderFactory) Create() wallets.WalletBuilder {
	out := createWalletBuilder()
	return out
}
