package objects

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
)

// Organization represents an organization
type Organization struct {
	Met      *concrete_metadata.MetaData `json:"metadata"`
	Creator  *Holder                     `json:"creator"`
	Shares   *Asset                      `json:"shares"`
	Quotas   *Asset                      `json:"quotas"`
	Currency *Currency                   `json:"currency"`
	Name     string                      `json:"name"`
	Desc     string                      `json:"description"`
}

// CreateOrganization creates a new Organization instance
func CreateOrganization(met metadata.MetaData, creator *Holder, shares *Asset, quotas *Asset, currency *Currency, name string, description string) *Organization {
	out := Organization{
		Met:      met.(*concrete_metadata.MetaData),
		Creator:  creator,
		Shares:   shares,
		Quotas:   quotas,
		Currency: currency,
		Name:     name,
		Desc:     description,
	}

	return &out
}
