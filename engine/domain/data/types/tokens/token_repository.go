package tokens

// TokenRepository represents a token repository
type TokenRepository interface {
	Retrieve(dirPath string) (Token, error)
}
