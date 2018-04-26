package objects

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
)

// Currency represents a currency
type Currency struct {
	Met     *concrete_metadata.MetaData `json:"metadata"`
	Creator *Holder                     `json:"creator"`
	Sym     string                      `json:"symbol"`
	Name    string                      `json:"name"`
	Desc    string                      `json:"description"`
}

// CreateCurrency creates a new currency instance
func CreateCurrency(met metadata.MetaData, creator *Holder, sym string, name string, description string) *Currency {
	out := Currency{
		Met:     met.(*concrete_metadata.MetaData),
		Creator: creator,
		Sym:     sym,
		Name:    name,
		Desc:    description,
	}

	return &out
}
