package wealth

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
)

// Token represents a token
type Token struct {
	Met     *metadata.MetaData `json:"metadata"`
	Creator *Entity            `json:"creator"`
	Symbol  string             `json:"symbol"`
	Amount  int                `json:"amount"`
}

// CreateToken returns the token instance
func CreateToken(met *metadata.MetaData, creator *Entity, symbol string, amount int) *Token {
	out := Token{
		Met:     met,
		Creator: creator,
		Symbol:  symbol,
		Amount:  amount,
	}

	return &out
}

// GetMetaData returns the metadata
func (tok *Token) GetMetaData() *metadata.MetaData {
	return tok.Met
}

// GetCreator returns the creator entity
func (tok *Token) GetCreator() *Entity {
	return tok.Creator
}

// GetSymbol returns the symbol
func (tok *Token) GetSymbol() string {
	return tok.Symbol
}

// GetAmount returns the amount
func (tok *Token) GetAmount() int {
	return tok.Amount
}
