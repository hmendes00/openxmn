package executors

import (
	"encoding/json"
	"log"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	executors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/executors"
	readdb "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	writedb "github.com/XMNBlockchain/openxmn/highway/project/databases/write"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Asset represents an insert asset executor
type Asset struct {
	readAssetDB  *readdb.Asset
	writeAssetDB *writedb.Asset
}

// CreateAsset creates a new Asset instance
func CreateAsset(readAssetDB *readdb.Asset, writeAssetDB *writedb.Asset) executors.Command {
	out := Asset{
		readAssetDB:  readAssetDB,
		writeAssetDB: writeAssetDB,
	}

	return &out
}

// Execute executes an asset executors.
func (db *Asset) Execute(cmd commands.Command) error {
	if cmd.HasInsert() {
		js := cmd.GetInsert().GetJS()
		newAsset := new(objects.Asset)
		jsErr := json.Unmarshal(js, newAsset)
		if jsErr != nil {
			return jsErr
		}

		insErr := db.writeAssetDB.Insert(newAsset)
		if insErr != nil {
			log.Printf("there was an error while inserting a new asset to the database: %s\n", insErr.Error())
			return nil
		}

		return nil
	}

	if cmd.HasDelete() {
		js := cmd.GetDelete().GetJS()
		delAsset := new(objects.Asset)
		jsErr := json.Unmarshal(js, delAsset)
		if jsErr != nil {
			return jsErr
		}

		delErr := db.writeAssetDB.Delete(delAsset)
		if delErr != nil {
			log.Printf("there was an error while deleting an existing asset from the database: %s\n", delErr.Error())
			return nil
		}

		return nil
	}

	if cmd.HasUpdate() {
		update := cmd.GetUpdate()
		originalJS := update.GetOriginalJS()
		originalAsset := new(objects.Asset)
		upJsErr := json.Unmarshal(originalJS, originalAsset)
		if upJsErr != nil {
			return upJsErr
		}

		newJS := update.GetNewJS()
		newAsset := new(objects.Asset)
		newJsErr := json.Unmarshal(newJS, newAsset)
		if newJsErr != nil {
			return newJsErr
		}

		upErr := db.writeAssetDB.Update(originalAsset, newAsset)
		if upErr != nil {
			log.Printf("there was an error while updating an existing asset from the database: %s\n", upErr.Error())
			return nil
		}

		return nil
	}

	return nil
}
