package objects

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
)

// AssetHolder represents an asset holder
type AssetHolder struct {
	Met    *concrete_metadata.MetaData `json:"metadata"`
	Usr    *Holder                     `json:"holder"`
	Ass    *Asset                      `json:"asset"`
	Amount uint                        `json:"amount"`
}

// CreateAssetHolder creates a new AssetHolder instance
func CreateAssetHolder(met metadata.MetaData, usr *Holder, ass *Asset, amount uint) *AssetHolder {
	out := AssetHolder{
		Met:    met.(*concrete_metadata.MetaData),
		Usr:    usr,
		Ass:    ass,
		Amount: amount,
	}

	return &out
}
