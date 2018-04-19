package commands

import (
	"path/filepath"

	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
)

// ChainedCommandsService represents a concrete chained commands service implementation
type ChainedCommandsService struct {
	metaDataService                     metadata.Service
	cmdsService                         commands.Service
	fileService                         files.FileService
	fileBuilderFactory                  files.FileBuilderFactory
	storedChainedCommandsBuilderFactory stored_commands.ChainedCommandsBuilderFactory
}

// CreateChainedCommandsService creates a new ChainedCommandsService instance
func CreateChainedCommandsService(metaDataService metadata.Service, cmdsService commands.Service, fileService files.FileService, fileBuilderFactory files.FileBuilderFactory, storedChainedCommandsBuilderFactory stored_commands.ChainedCommandsBuilderFactory) commands.ChainedCommandsService {
	out := ChainedCommandsService{
		metaDataService:                     metaDataService,
		cmdsService:                         cmdsService,
		fileService:                         fileService,
		fileBuilderFactory:                  fileBuilderFactory,
		storedChainedCommandsBuilderFactory: storedChainedCommandsBuilderFactory,
	}

	return &out
}

// Save saves a ChainedCommands instance
func (serv *ChainedCommandsService) Save(dirPath string, cmd commands.ChainedCommands) (stored_commands.ChainedCommands, error) {
	met := cmd.GetMetaData()
	metFile, metFileErr := serv.metaDataService.Save(dirPath, met)
	if metFileErr != nil {
		return nil, metFileErr
	}

	cmds := cmd.GetCommands()
	cmdsPath := filepath.Join(dirPath, "commands")
	storedCmds, storedCmdsErr := serv.cmdsService.Save(cmdsPath, cmds)
	if storedCmdsErr != nil {
		return nil, storedCmdsErr
	}

	previousIDAsString := cmd.GetPreviousID().String()
	prevIDFile, prevIDFileErr := serv.fileBuilderFactory.Create().Create().WithData([]byte(previousIDAsString)).WithFileName("previous").WithExtension("id").Now()
	if prevIDFileErr != nil {
		return nil, prevIDFileErr
	}

	prevIDStoredFile, prevIDStoredFileERR := serv.fileService.Save(dirPath, prevIDFile)
	if prevIDStoredFileERR != nil {
		return nil, prevIDStoredFileERR
	}

	rootAsString := cmd.GetRootID().String()
	rootIDFile, rootIDFileErr := serv.fileBuilderFactory.Create().Create().WithData([]byte(rootAsString)).WithFileName("root").WithExtension("id").Now()
	if rootIDFileErr != nil {
		return nil, rootIDFileErr
	}

	rootIDStoredFile, rootIDStoredFileErr := serv.fileService.Save(dirPath, rootIDFile)
	if rootIDStoredFileErr != nil {
		return nil, rootIDStoredFileErr
	}

	out, outErr := serv.storedChainedCommandsBuilderFactory.Create().Create().WithMetaData(metFile).WithCommands(storedCmds).WithPreviousID(prevIDStoredFile).WithRootID(rootIDStoredFile).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
