package blockchains

// BlockchainRepository represents a blockchain repository
type BlockchainRepository interface {
	Retrieve(dirPath string) (Blockchain, error)
}
