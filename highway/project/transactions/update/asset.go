package update

import uuid "github.com/satori/go.uuid"

// Asset represents an update asset transaction
type Asset struct {
	AssetID *uuid.UUID `json:"asset_id"`
	Sym     string     `json:"symbol"`
	Name    string     `json:"name"`
	Desc    string     `json:"description"`
}
