package executors

import (
	"encoding/json"
	"log"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	executors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/executors"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	writedb "github.com/XMNBlockchain/openxmn/highway/project/databases/write"
)

// User represents a user executor
type User struct {
	writeDB *writedb.User
}

// CreateUser creates a new User instance
func CreateUser(writeDB *writedb.User) executors.Command {
	out := User{
		writeDB: writeDB,
	}

	return &out
}

// Execute executes an asset executors.
func (db *User) Execute(cmd commands.Command) error {
	if cmd.HasInsert() {
		js := cmd.GetInsert().GetJS()
		newUser := new(concrete_users.User)
		jsErr := json.Unmarshal(js, newUser)
		if jsErr != nil {
			return jsErr
		}

		db.writeDB.Insert(newUser)
		return nil
	}

	if cmd.HasUpdate() {
		update := cmd.GetUpdate()
		originalJS := update.GetOriginalJS()
		originalUser := new(concrete_users.User)
		originalJSErr := json.Unmarshal(originalJS, originalUser)
		if originalJSErr != nil {
			return originalJSErr
		}

		newJS := update.GetNewJS()
		newUser := new(concrete_users.User)
		newJSErr := json.Unmarshal(newJS, newUser)
		if newJSErr != nil {
			return newJSErr
		}

		upUserErr := db.writeDB.Update(originalUser, newUser)
		if upUserErr != nil {
			log.Printf("there was an error while updating an existing user from the database: %s\n", upUserErr.Error())
			return nil
		}

		return nil
	}

	if cmd.HasDelete() {
		delJS := cmd.GetDelete().GetJS()
		delUser := new(concrete_users.User)
		delUserJSErr := json.Unmarshal(delJS, delUser)
		if delUserJSErr != nil {
			return delUserJSErr
		}

		delUserErr := db.writeDB.Delete(delUser)
		if delUserErr != nil {
			log.Printf("there was an error while deleting an existing user from the database: %s\n", delUserErr.Error())
			return nil
		}

		return nil
	}

	return nil
}
