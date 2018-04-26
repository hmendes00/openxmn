package write

import (
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Asset represents an asset write database
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

// Insert inserts a new asset
func (db *Asset) Insert(ass *objects.Asset) error {
	return nil
}

// Update updates an existing asset
func (db *Asset) Update(original *objects.Asset, new *objects.Asset) error {
	return nil
}

// Delete deletes an existing asset
func (db *Asset) Delete(ass *objects.Asset) error {
	return nil
}
