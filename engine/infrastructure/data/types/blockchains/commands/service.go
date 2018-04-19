package commands

import (
	"path/filepath"

	stored_chained_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/chunks"
)

// Service represents a concrete commands service
type Service struct {
	metaDataService          metadata.Service
	chksService              chunks.Service
	chksBuilderFactory       chunks.BuilderFactory
	storedCmdsBuilderFactory stored_chained_commands.BuilderFactory
}

// CreateService creates a new Service instance
func CreateService(metaDataService metadata.Service, chksService chunks.Service, chksBuilderFactory chunks.BuilderFactory, storedCmdsBuilderFactory stored_chained_commands.BuilderFactory) commands.Service {
	out := Service{
		metaDataService:          metaDataService,
		chksService:              chksService,
		chksBuilderFactory:       chksBuilderFactory,
		storedCmdsBuilderFactory: storedCmdsBuilderFactory,
	}

	return &out
}

// Save saves a Commands instance
func (serv *Service) Save(dirPath string, cmds commands.Commands) (stored_chained_commands.Commands, error) {
	met := cmds.GetMetaData()
	metFile, metFileErr := serv.metaDataService.Save(dirPath, met)
	if metFileErr != nil {
		return nil, metFileErr
	}

	chks := []chunks.Chunks{}
	cmdsList := cmds.GetCommands()
	for _, oneCmd := range cmdsList {
		chk, chkErr := serv.chksBuilderFactory.Create().Create().WithInstance(oneCmd).Now()
		if chkErr != nil {
			return nil, chkErr
		}

		chks = append(chks, chk)
	}

	commandsPath := filepath.Join(dirPath, "commands")
	chkFiles, chkFilesErr := serv.chksService.SaveAll(commandsPath, chks)
	if chkFilesErr != nil {
		return nil, chkFilesErr
	}

	out, outErr := serv.storedCmdsBuilderFactory.Create().Create().WithMetaData(metFile).WithCommands(chkFiles).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
