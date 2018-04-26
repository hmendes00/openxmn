package other

import (
	uuid "github.com/satori/go.uuid"
)

// TransferAsset represents a transfer asset transaction
type TransferAsset struct {
	FromOrgID *uuid.UUID `json:"from_organization_id"`
	ToUserID  *uuid.UUID `json:"to_user_id"`
	ToOrgID   *uuid.UUID `json:"to_organization_id"`
	AssetID   *uuid.UUID `json:"asset_id"`
	Amount    int        `json:"amount"`
}
