package executors

import (
	"encoding/json"
	"log"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	executors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/executors"
	writedb "github.com/XMNBlockchain/openxmn/highway/project/databases/write"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Project represents a project executor
type Project struct {
	writeDB *writedb.Project
}

// CreateProject creates a new Project instance
func CreateProject(writeDB *writedb.Project) executors.Command {
	out := Project{
		writeDB: writeDB,
	}

	return &out
}

// Execute executes an asset executors.
func (db *Project) Execute(cmd commands.Command) error {
	if cmd.HasInsert() {
		js := cmd.GetInsert().GetJS()
		newProject := new(objects.Project)
		jsErr := json.Unmarshal(js, newProject)
		if jsErr != nil {
			return jsErr
		}

		insProjErr := db.writeDB.Insert(newProject)
		if insProjErr != nil {
			log.Printf("there was an error while inserting a new project to the database: %s\n", insProjErr.Error())
			return nil
		}

		return nil
	}

	if cmd.HasUpdate() {
		update := cmd.GetUpdate()
		originalJS := update.GetOriginalJS()
		originalProject := new(objects.Project)
		originalJSErr := json.Unmarshal(originalJS, originalProject)
		if originalJSErr != nil {
			return originalJSErr
		}

		newJS := update.GetNewJS()
		newProject := new(objects.Project)
		newJSErr := json.Unmarshal(newJS, newProject)
		if newJSErr != nil {
			return newJSErr
		}

		upProjErr := db.writeDB.Update(originalProject, newProject)
		if upProjErr != nil {
			log.Printf("there was an error while updating an existing project from the database: %s\n", upProjErr.Error())
			return nil
		}

		return nil
	}

	if cmd.HasDelete() {
		delJS := cmd.GetDelete().GetJS()
		delProject := new(objects.Project)
		delProjJSErr := json.Unmarshal(delJS, delProject)
		if delProjJSErr != nil {
			return delProjJSErr
		}

		delProjectErr := db.writeDB.Delete(delProject)
		if delProjectErr != nil {
			log.Printf("there was an error while deleting an existing project from the database: %s\n", delProjectErr.Error())
			return nil
		}

		return nil
	}

	return nil
}
