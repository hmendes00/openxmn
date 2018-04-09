package databases

import (
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	uuid "github.com/satori/go.uuid"
)

// Token represents a token database
type Token struct {
	dirPath    string
	repository tokens.TokenRepository
}

// CreateToken creates a new Token database
func CreateToken(dirPath string, repository tokens.TokenRepository) *Token {
	out := Token{
		dirPath:    dirPath,
		repository: repository,
	}

	return &out
}

// RetrieveByID retrieves a Token by ID
func (db *Token) RetrieveByID(id *uuid.UUID) (tokens.Token, error) {
	return nil, nil
}
