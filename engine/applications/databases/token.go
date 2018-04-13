package databases

import (
	"errors"
	"fmt"

	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	uuid "github.com/satori/go.uuid"
)

// Token represents a token database
type Token struct {
	toks map[string]tokens.Token
}

// CreateToken creates a new Token database
func CreateToken(dirPath string) *Token {
	out := Token{
		toks: map[string]tokens.Token{},
	}

	return &out
}

// RetrieveByID retrieves a Token by ID
func (db *Token) RetrieveByID(id *uuid.UUID) (tokens.Token, error) {
	idAsString := id.String()
	if oneTok, ok := db.toks[idAsString]; ok {
		return oneTok, nil
	}

	str := fmt.Sprintf("the token (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

// Insert inserts a new token to the database
func (db *Token) Insert(tok tokens.Token) error {
	id := tok.GetMetaData().GetID()
	idAsString := id.String()
	_, retTokErr := db.RetrieveByID(id)
	if retTokErr == nil {
		str := fmt.Sprintf("there is already a token with ID: %s", idAsString)
		return errors.New(str)
	}

	db.toks[idAsString] = tok
	return nil
}

// Update updates a token to the database
func (db *Token) Update(old tokens.Token, new tokens.Token) error {
	newTokID := new.GetMetaData().GetID()
	newTokIDAsString := newTokID.String()
	_, retNewTokErr := db.RetrieveByID(newTokID)
	if retNewTokErr == nil {
		str := fmt.Sprintf("the new token (ID: %s) already exists", newTokIDAsString)
		return errors.New(str)
	}

	delErr := db.Delete(old)
	if delErr != nil {
		return delErr
	}

	insErr := db.Insert(new)
	if insErr != nil {
		return insErr
	}

	return nil
}

// Delete deletes a token
func (db *Token) Delete(tok tokens.Token) error {
	id := tok.GetMetaData().GetID()
	_, retTokErr := db.RetrieveByID(id)
	if retTokErr != nil {
		return retTokErr
	}

	idAsString := id.String()
	delete(db.toks, idAsString)
	return nil
}
