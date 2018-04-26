package insert

import uuid "github.com/satori/go.uuid"

// Asset represents an insert asset transaction
type Asset struct {
	AssetID *uuid.UUID `json:"asset_id"`
	CrOrgID *uuid.UUID `json:"creator_organization_id"`
	Sym     string     `json:"symbol"`
	Name    string     `json:"name"`
	Desc    string     `json:"description"`
	Amount  uint       `json:"amount"`
}
