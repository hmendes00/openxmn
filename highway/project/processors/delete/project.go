package delete

import (
	"encoding/json"
	"errors"
	"fmt"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	transaction_delete "github.com/XMNBlockchain/openxmn/highway/project/transactions/delete"
)

// Project represents a delete project processor
type Project struct {
	projectDB            *database.Project
	deleteBuilderFactory commands.DeleteBuilderFactory
	cmdBuilderFactory    commands.CommandBuilderFactory
}

// CreateProject creates a new Project instance
func CreateProject(
	projectDB *database.Project,
	deleteBuilderFactory commands.DeleteBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Project{
		projectDB:            projectDB,
		deleteBuilderFactory: deleteBuilderFactory,
		cmdBuilderFactory:    cmdBuilderFactory,
	}

	return &out
}

// Process processes a Project transaction
func (proc *Project) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	projTrs := new(transaction_delete.Project)
	jsErr := json.Unmarshal(js, projTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//retrieve the project:
	proj, projErr := proc.projectDB.RetrieveByID(projTrs.ProjectID)
	if projErr != nil {
		return nil, projErr
	}

	//make sure the user has the right to update the project:
	if !proc.projectDB.CanDelete(proj, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to delete the project (ID: %s)", user.GetMetaData().GetID().String(), proj.Met.GetID().String())
		return nil, errors.New(str)
	}

	//convert the project to JS:
	projJS, projJSErr := json.Marshal(proj)
	if projJSErr != nil {
		return nil, projJSErr
	}

	//build the delete command:
	del, delErr := proc.deleteBuilderFactory.Create().Create().WithJS(projJS).Now()
	if delErr != nil {
		return nil, delErr
	}

	//build the command:
	cmd, cmdErr := proc.cmdBuilderFactory.Create().Create().WithDelete(del).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
