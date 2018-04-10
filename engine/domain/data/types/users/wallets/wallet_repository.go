package wallets

// WalletRepository represents a wallet repository
type WalletRepository interface {
	Retrieve(dirPath string) (Wallet, error)
}
