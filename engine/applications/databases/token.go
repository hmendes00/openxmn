package databases

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
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

// Insert inserts a new token to the database
func (db *Token) Insert(tok tokens.Token) (files.File, error) {
	return nil, nil
}

// Update updates a token to the database
func (db *Token) Update(old tokens.Token, new tokens.Token) (files.File, files.File, error) {
	return nil, nil, nil
}

// Delete deletes a token
func (db *Token) Delete(tok tokens.Token) (files.File, error) {
	return nil, nil
}
