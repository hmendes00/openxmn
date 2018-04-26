package update

import (
	"encoding/json"
	"errors"
	"fmt"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	transaction_update "github.com/XMNBlockchain/openxmn/highway/project/transactions/update"
)

// Project represents an update project processor
type Project struct {
	projectDB              *database.Project
	metaDataBuilderFactory metadata.BuilderFactory
	updateBuilderFactory   commands.UpdateBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateProject creates a new Project instance
func CreateProject(
	projectDB *database.Project,
	metaDataBuilderFactory metadata.BuilderFactory,
	updateBuilderFactory commands.UpdateBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Project{
		projectDB:              projectDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		updateBuilderFactory:   updateBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes a Project transaction
func (proc *Project) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	projTrs := new(transaction_update.Project)
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
	if !proc.projectDB.CanUpdate(proj, user) {
		str := fmt.Sprintf("the user (ID: %s) do not have the right to update the project (ID: %s)", user.GetMetaData().GetID().String(), proj.Met.GetID().String())
		return nil, errors.New(str)
	}

	//build the metadata:
	id := proj.Met.GetID()
	crOn := proj.Met.CreatedOn()
	lstOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(crOn).LastUpdatedOn(lstOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//create the new project:
	newProject, newProjectErr := objects.CreateProject(
		met,
		proj.RelPath,
		proj.Org,
		projTrs.PricePerTrx,
		projTrs.ShareToProcessors,
		projTrs.ShareToLeaders,
		projTrs.ShareToVerifiers,
		projTrs.ShareToBlocker,
		projTrs.ShareToShareHolders,
		projTrs.AmountOfQuotasNeededPerBlock,
		projTrs.BlockDuration,
	)

	if newProjectErr != nil {
		return nil, newProjectErr
	}

	//convert the project to JS:
	projJS, projJSErr := json.Marshal(newProject)
	if projJSErr != nil {
		return nil, projJSErr
	}

	//convert the original project to JS:
	originalProjJS, originalProjJSErr := json.Marshal(proj)
	if originalProjJSErr != nil {
		return nil, originalProjJSErr
	}

	//build the update command:
	up, upErr := proc.updateBuilderFactory.Create().Create().WithNewJS(projJS).WithOriginalJS(originalProjJS).Now()
	if upErr != nil {
		return nil, upErr
	}

	//build the command:
	cmd, cmdErr := proc.cmdBuilderFactory.Create().Create().WithUpdate(up).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
