package objects

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
)

// CurrencyHolder represents a currency holder
type CurrencyHolder struct {
	Met    *concrete_metadata.MetaData `json:"metadata"`
	Usr    *Holder                     `json:"holder"`
	Curr   *Currency                   `json:"currency"`
	Amount float64                     `json:"amount"`
}

// CreateCurrencyHolder creates a new CurrencyHolder instance
func CreateCurrencyHolder(met metadata.MetaData, usr *Holder, curr *Currency, amount float64) *CurrencyHolder {
	out := CurrencyHolder{
		Met:    met.(*concrete_metadata.MetaData),
		Usr:    usr,
		Curr:   curr,
		Amount: amount,
	}

	return &out
}
