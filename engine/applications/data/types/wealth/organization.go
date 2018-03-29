package wealth

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
	user "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Organization represents an organized group of user
type Organization struct {
	Met                       *metadata.MetaData `json:"metadata"`
	Usr                       *user.User         `json:"user"`
	AcceptedToken             *Token             `json:"accepted_token"`
	PercentNeededForConcensus float64            `json:"percent_needed_for_concensus"`
}

// CreateOrganization creates a new Organization instance
func CreateOrganization(met *metadata.MetaData, usr *user.User, acceptedToken *Token, percentNeeded float64) *Organization {
	out := Organization{
		Met:                       met,
		Usr:                       usr,
		AcceptedToken:             acceptedToken,
		PercentNeededForConcensus: percentNeeded,
	}

	return &out
}

// GetMetaData returns the metadata
func (org *Organization) GetMetaData() *metadata.MetaData {
	return org.Met
}

// CreatedByUser returns the user creator
func (org *Organization) CreatedByUser() *user.User {
	return org.Usr
}

// GetAcceptedToken returns the accepted token
func (org *Organization) GetAcceptedToken() *Token {
	return org.AcceptedToken
}

// GetPercentNeededForConcensus returns the minimum percentage needed to reach a concensus
func (org *Organization) GetPercentNeededForConcensus() float64 {
	return org.PercentNeededForConcensus
}
