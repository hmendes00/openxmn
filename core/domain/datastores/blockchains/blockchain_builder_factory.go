package blockchains

// BlockchainBuilderFactory represents a blockchain builder factory
type BlockchainBuilderFactory interface {
	Create() BlockchainBuilder
}
