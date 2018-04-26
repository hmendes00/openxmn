package read

import (
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	uuid "github.com/satori/go.uuid"
)

// Asset represents an asset read database
type Asset struct {
	assets map[string]*objects.Asset
}

// CreateAsset creates a new Asset instance
func CreateAsset(assets map[string]*objects.Asset) *Asset {
	out := Asset{
		assets: assets,
	}

	return &out
}

// RetrieveByID retrieves an asset by its ID
func (db *Asset) RetrieveByID(id *uuid.UUID) (*objects.Asset, error) {
	return nil, nil
}

// CanUpdate verifies if a given user can update the given asset
func (db *Asset) CanUpdate(asset *objects.Asset, user users.User) bool {
	return true
}

// CanDelete verifies if a given user can delete the given asset
func (db *Asset) CanDelete(asset *objects.Asset, user users.User) bool {
	return true
}
