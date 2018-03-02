package domain

// ChainRepository represents the chain repository
type ChainRepository interface {
	Retrieve(dirPath string) (Chain, error)
}
