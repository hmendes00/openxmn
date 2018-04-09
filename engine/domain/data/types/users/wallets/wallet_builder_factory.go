package wallets

// WalletBuilderFactory represents a wallet builder factory
type WalletBuilderFactory interface {
	Create() WalletBuilder
}
