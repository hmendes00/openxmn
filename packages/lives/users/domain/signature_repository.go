package domain

// SignatureRepository represents a Signature repository
type SignatureRepository interface {
	Retrieve(dirPath string) (Signature, error)
	RetrieveAll(dirPath string) ([]Signature, error)
}
