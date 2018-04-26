package delete

import uuid "github.com/satori/go.uuid"

// Asset represents a delete asset transaction
type Asset struct {
	AssetID *uuid.UUID `json:"asset_id"`
}
