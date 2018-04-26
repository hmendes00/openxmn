package insert

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	database "github.com/XMNBlockchain/openxmn/highway/project/databases/read"
	objects "github.com/XMNBlockchain/openxmn/highway/project/objects"
	transaction_insert "github.com/XMNBlockchain/openxmn/highway/project/transactions/insert"
)

// Project represents an insert project processor
type Project struct {
	projectDB              *database.Project
	orgDB                  *database.Organization
	metaDataBuilderFactory metadata.BuilderFactory
	insertBuilderFactory   commands.InsertBuilderFactory
	cmdBuilderFactory      commands.CommandBuilderFactory
}

// CreateProject creates a new Project instance
func CreateProject(
	projectDB *database.Project,
	orgDB *database.Organization,
	metaDataBuilderFactory metadata.BuilderFactory,
	insertBuilderFactory commands.InsertBuilderFactory,
	cmdBuilderFactory commands.CommandBuilderFactory,
) processors.Transaction {
	out := Project{
		projectDB: projectDB,
		orgDB:     orgDB,
		metaDataBuilderFactory: metaDataBuilderFactory,
		insertBuilderFactory:   insertBuilderFactory,
		cmdBuilderFactory:      cmdBuilderFactory,
	}

	return &out
}

// Process processes a Project transaction
func (proc *Project) Process(trs transactions.Transaction, user users.User) (commands.Command, error) {
	//try to unmarshal:
	js := trs.GetJSON()
	projTrs := new(transaction_insert.Project)
	jsErr := json.Unmarshal(js, projTrs)
	if jsErr != nil {
		return nil, jsErr
	}

	//make sure the project does not already exists:
	_, projErr := proc.projectDB.RetrieveByID(projTrs.ProjectID)
	if projErr == nil {
		str := fmt.Sprintf("the project (ID: %s) already exists", projTrs.ProjectID.String())
		return nil, errors.New(str)
	}

	//retrieve the organization:
	org, orgErr := proc.orgDB.RetrieveByID(projTrs.OrgID)
	if orgErr != nil {
		return nil, orgErr
	}

	//build the metadata:
	crOn := trs.GetMetaData().CreatedOn()
	met, metErr := proc.metaDataBuilderFactory.Create().Create().WithID(projTrs.ProjectID).CreatedOn(crOn).Now()
	if metErr != nil {
		return nil, metErr
	}

	//create the relative path:
	relPath := path.Join(
		string(crOn.Year()),
		string(crOn.Month()),
		string(crOn.Day()),
		string(crOn.Hour()),
		string(crOn.Minute()),
		string(crOn.Second()),
		projTrs.ProjectID.String(),
	)

	//create the new project:
	newProject, newProjectErr := objects.CreateProject(
		met,
		relPath,
		org,
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

	//build the insert command:
	ins, insErr := proc.insertBuilderFactory.Create().Create().WithJS(projJS).Now()
	if insErr != nil {
		return nil, insErr
	}

	//build the command:
	cmd, cmdErr := proc.cmdBuilderFactory.Create().Create().WithInsert(ins).Now()
	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, nil
}
