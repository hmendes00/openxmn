package executors

import (
	"encoding/json"
	"log"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	executors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/executors"
	writedb "github.com/XMNBlockchain/openxmn/highway/project/databases/write"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Organization represents a organization executor
type Organization struct {
	writeDB *writedb.Organization
}

// CreateOrganization creates a new Organization instance
func CreateOrganization(writeDB *writedb.Organization) executors.Command {
	out := Organization{
		writeDB: writeDB,
	}

	return &out
}

// Execute executes an asset executors.
func (db *Organization) Execute(cmd commands.Command) error {
	if cmd.HasInsert() {
		js := cmd.GetInsert().GetJS()
		newOrganization := new(objects.Organization)
		jsErr := json.Unmarshal(js, newOrganization)
		if jsErr != nil {
			return jsErr
		}

		db.writeDB.Insert(newOrganization)
		return nil
	}

	if cmd.HasUpdate() {
		update := cmd.GetUpdate()
		originalJS := update.GetOriginalJS()
		originalOrganization := new(objects.Organization)
		originalJSErr := json.Unmarshal(originalJS, originalOrganization)
		if originalJSErr != nil {
			return originalJSErr
		}

		newJS := update.GetNewJS()
		newOrganization := new(objects.Organization)
		newJSErr := json.Unmarshal(newJS, newOrganization)
		if newJSErr != nil {
			return newJSErr
		}

		upOrgErr := db.writeDB.Update(originalOrganization, newOrganization)
		if upOrgErr != nil {
			log.Printf("there was an error while updating an existing organization from the database: %s\n", upOrgErr.Error())
			return nil
		}

		return nil
	}

	if cmd.HasDelete() {
		delJS := cmd.GetDelete().GetJS()
		delOrganization := new(objects.Organization)
		delOrgJSErr := json.Unmarshal(delJS, delOrganization)
		if delOrgJSErr != nil {
			return delOrgJSErr
		}

		delOrganizationErr := db.writeDB.Delete(delOrganization)
		if delOrganizationErr != nil {
			log.Printf("there was an error while deleting an existing organization from the database: %s\n", delOrganizationErr.Error())
			return nil
		}

		return nil
	}

	return nil
}
