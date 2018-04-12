package tokens

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

// Token represents a concrete token implementation
type Token struct {
	Met     *concrete_metadata.MetaData `json:"metadata"`
	Creator *concrete_users.User        `json:"creator"`
	Symbol  string                      `json:"symbol"`
}

func createToken(met *concrete_metadata.MetaData, creator *concrete_users.User, symbol string) tokens.Token {
	out := Token{
		Met:     met,
		Creator: creator,
		Symbol:  symbol,
	}

	return &out
}

// GetMetaData returns the metadata
func (tok *Token) GetMetaData() metadata.MetaData {
	return tok.Met
}

// GetCreator returns the creator
func (tok *Token) GetCreator() users.User {
	return tok.Creator
}

// GetSymbol returns the symbol
func (tok *Token) GetSymbol() string {
	return tok.Symbol
}
