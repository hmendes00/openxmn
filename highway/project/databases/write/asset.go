package write

import (
	"errors"
	"fmt"

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
func (db *Asset) Insert(ass *objects.Asset) {
	db.assets[ass.Met.GetID().String()] = ass
}

// Update updates an existing asset
func (db *Asset) Update(original *objects.Asset, new *objects.Asset) error {
	delErr := db.Delete(original)
	if delErr != nil {
		return delErr
	}

	db.Insert(new)
	return nil
}

// Delete deletes an existing asset
func (db *Asset) Delete(ass *objects.Asset) error {
	idAsString := ass.Met.GetID().String()
	if _, ok := db.assets[idAsString]; ok {
		delete(db.assets, idAsString)
		return nil
	}

	str := fmt.Sprintf("the asset (ID: %s) could not be found", idAsString)
	return errors.New(str)
}
