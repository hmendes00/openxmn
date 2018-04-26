package other

import (
	uuid "github.com/satori/go.uuid"
)

// SplitAsset represents a split asset transaction
type SplitAsset struct {
	AssetID *uuid.UUID `json:"asset_id"`
	Amount  uint       `json:"amount"`
}
