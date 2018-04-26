package executors

import (
	"encoding/json"
	"log"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	executors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/executors"
	writedb "github.com/XMNBlockchain/openxmn/highway/project/databases/write"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
)

// Server represents a server executor
type Server struct {
	writeDB *writedb.Server
}

// CreateServer creates a new Server instance
func CreateServer(writeDB *writedb.Server) executors.Command {
	out := Server{
		writeDB: writeDB,
	}

	return &out
}

// Execute executes an asset executors.
func (db *Server) Execute(cmd commands.Command) error {
	if cmd.HasInsert() {
		js := cmd.GetInsert().GetJS()
		newServer := new(objects.Server)
		jsErr := json.Unmarshal(js, newServer)
		if jsErr != nil {
			return jsErr
		}

		insServErr := db.writeDB.Insert(newServer)
		if insServErr != nil {
			log.Printf("there was an error while inserting a new server to the database: %s\n", insServErr.Error())
			return nil
		}

		return nil
	}

	if cmd.HasUpdate() {
		update := cmd.GetUpdate()
		originalJS := update.GetOriginalJS()
		originalServer := new(objects.Server)
		originalJSErr := json.Unmarshal(originalJS, originalServer)
		if originalJSErr != nil {
			return originalJSErr
		}

		newJS := update.GetNewJS()
		newServer := new(objects.Server)
		newJSErr := json.Unmarshal(newJS, newServer)
		if newJSErr != nil {
			return newJSErr
		}

		upServErr := db.writeDB.Update(originalServer, newServer)
		if upServErr != nil {
			log.Printf("there was an error while updating an existing server from the database: %s\n", upServErr.Error())
			return nil
		}

		return nil
	}

	if cmd.HasDelete() {
		delJS := cmd.GetDelete().GetJS()
		delServer := new(objects.Server)
		delServJSErr := json.Unmarshal(delJS, delServer)
		if delServJSErr != nil {
			return delServJSErr
		}

		delServerErr := db.writeDB.Delete(delServer)
		if delServerErr != nil {
			log.Printf("there was an error while deleting an existing server from the database: %s\n", delServerErr.Error())
			return nil
		}

		return nil
	}

	return nil
}
