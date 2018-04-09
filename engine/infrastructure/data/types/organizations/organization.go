package organizations

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

// Organization represents a concrete organization implementation
type Organization struct {
	Met                       *concrete_metadata.MetaData `json:"metadata"`
	Usr                       *concrete_users.User        `json:"user"`
	Tok                       *concrete_tokens.Token      `json:"accepted_token"`
	PercentNeededForConcensus float64                     `json:"percent_needed_for_concensus"`
}

func createOrganization(met *concrete_metadata.MetaData, usr *concrete_users.User, tok *concrete_tokens.Token, percentNeededForConcensus float64) organizations.Organization {
	out := Organization{
		Met: met,
		Usr: usr,
		Tok: tok,
		PercentNeededForConcensus: percentNeededForConcensus,
	}

	return &out
}

// GetMetaData returns the metadata
func (org *Organization) GetMetaData() metadata.MetaData {
	return org.Met
}

// GetUser returns the user
func (org *Organization) GetUser() users.User {
	return org.Usr
}

// GetAcceptedToken returns the accepted token
func (org *Organization) GetAcceptedToken() tokens.Token {
	return org.Tok
}

// GetPercentNeededForConcensus returns the percent needed for concensus
func (org *Organization) GetPercentNeededForConcensus() float64 {
	return org.PercentNeededForConcensus
}
