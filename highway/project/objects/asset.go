package objects

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
)

// Asset represents an organization asset
type Asset struct {
	Met     *concrete_metadata.MetaData `json:"metadata"`
	Creator *Holder                     `json:"creator"`
	Sym     string                      `json:"symbol"`
	Name    string                      `json:"name"`
	Desc    string                      `json:"description"`
}

// CreateAsset creates a new Asset instance
func CreateAsset(met metadata.MetaData, creator *Holder, sym string, name string, desc string) *Asset {
	out := Asset{
		Met:     met.(*concrete_metadata.MetaData),
		Creator: creator,
		Sym:     sym,
		Name:    name,
		Desc:    desc,
	}

	return &out
}
