package blockchains

// BlockchainService represents a blockchain service
type BlockchainService interface {
	Chain(dirPath string, validatedBlk string) (Blockchain, error)
}
