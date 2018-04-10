package wallets

import (
	"github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

// WalletBuilderFactory represents a concrete WalletBuilderFactory implementation
type WalletBuilderFactory struct {
	metaDataBuilderFactory metadata.BuilderFactory
}

// CreateWalletBuilderFactory creates a new WalletBuilderFactory implementation
func CreateWalletBuilderFactory(metaDataBuilderFactory metadata.BuilderFactory) wallets.WalletBuilderFactory {
	out := WalletBuilderFactory{
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new WalletBuilder instance
func (fac *WalletBuilderFactory) Create() wallets.WalletBuilder {
	out := createWalletBuilder(fac.metaDataBuilderFactory)
	return out
}
