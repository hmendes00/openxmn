package commands

import (
	"path/filepath"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
	uuid "github.com/satori/go.uuid"
)

// ChainedCommandsRepository represents a concrete chained commands repository implementation
type ChainedCommandsRepository struct {
	metaDataRepository metadata.Repository
	cmdsRepository     commands.Repository
	fileRepository     files.FileRepository
	cmdsBuilderFactory commands.ChainedCommandsBuilderFactory
}

// CreateChainedCommandsRepository creates a ChainedCommandsRepository instance
func CreateChainedCommandsRepository(metaDataRepository metadata.Repository, cmdsRepository commands.Repository, fileRepository files.FileRepository, cmdsBuilderFactory commands.ChainedCommandsBuilderFactory) commands.ChainedCommandsRepository {
	out := ChainedCommandsRepository{
		metaDataRepository: metaDataRepository,
		cmdsRepository:     cmdsRepository,
		fileRepository:     fileRepository,
		cmdsBuilderFactory: cmdsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a ChainedCommands instance
func (rep *ChainedCommandsRepository) Retrieve(dirPath string) (commands.ChainedCommands, error) {
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	commandsPath := filepath.Join(dirPath, "commands")
	cmds, cmdsErr := rep.cmdsRepository.Retrieve(commandsPath)
	if cmdsErr != nil {
		return nil, cmdsErr
	}

	prevIDFile, prevIDFileErr := rep.fileRepository.Retrieve(dirPath, "previous.id")
	if prevIDFileErr != nil {
		return nil, prevIDFileErr
	}

	prevIDAsString := string(prevIDFile.GetData())
	prevID, prevIDErr := uuid.FromString(prevIDAsString)
	if prevIDErr != nil {
		return nil, prevIDErr
	}

	rootIDFile, rootIDFileErr := rep.fileRepository.Retrieve(dirPath, "root.id")
	if rootIDFileErr != nil {
		return nil, rootIDFileErr
	}

	rootIDAsString := string(rootIDFile.GetData())
	rootID, rootIDErr := uuid.FromString(rootIDAsString)
	if rootIDErr != nil {
		return nil, rootIDErr
	}

	out, outErr := rep.cmdsBuilderFactory.Create().Create().WithMetaData(met).WithCommands(cmds).WithPreviousID(&prevID).WithRootID(&rootID).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
