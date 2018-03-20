package transactions

import (
	uuid "github.com/satori/go.uuid"
)

// SaveOrganization represents a save organization transaction
type SaveOrganization struct {
	ID                        *uuid.UUID `json:"id"`
	TokenID                   *uuid.UUID `json:"token_id"`
	PercentNeededForConcensus float64    `json:"percent_needed_for_concensus"`
}
