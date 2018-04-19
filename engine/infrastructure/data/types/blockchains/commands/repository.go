package commands

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/chunks"
)

// Repository represents a concrete commands repository
type Repository struct {
	metaDataRepository metadata.Repository
	chunksRepository   chunks.Repository
	cmdsBuilderFactory commands.BuilderFactory
}

// CreateRepository creates a new Repository instance
func CreateRepository(metaDataRepository metadata.Repository, chunksRepository chunks.Repository, cmdsBuilderFactory commands.BuilderFactory) commands.Repository {
	out := Repository{
		metaDataRepository: metaDataRepository,
		chunksRepository:   chunksRepository,
		cmdsBuilderFactory: cmdsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a new Commands instance
func (rep *Repository) Retrieve(dirPath string) (commands.Commands, error) {
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the commands chunks:
	commandsPath := filepath.Join(dirPath, "commands")
	chks, chksErr := rep.chunksRepository.RetrieveAll(commandsPath)
	if chksErr != nil {
		return nil, chksErr
	}

	commandsList := []commands.Command{}
	for _, oneChk := range chks {
		oneCmd := new(Command)
		cmdErr := oneChk.Marshal(oneCmd)
		if cmdErr != nil {
			return nil, cmdErr
		}

		commandsList = append(commandsList, oneCmd)
	}

	blocks := [][]byte{
		met.GetID().Bytes(),
		[]byte(strconv.Itoa(int(met.CreatedOn().UnixNano()))),
	}

	cmdMap := map[string]commands.Command{}
	for _, oneCmd := range commandsList {
		hash := oneCmd.GetMetaData().GetHashTree().GetHash()
		cmdMap[hash.String()] = oneCmd
		blocks = append(blocks, hash.Get())
	}

	//re-order the blocks:
	reOrderedBlks, reOrderedBlksErr := met.GetHashTree().Order(blocks)
	if reOrderedBlksErr != nil {
		return nil, reOrderedBlksErr
	}

	//re-order the commands:
	reOrderedCmds := []commands.Command{}
	for _, oneBlk := range reOrderedBlks[2:] {
		blkAsString := hex.EncodeToString(oneBlk)
		if oneCmd, ok := cmdMap[blkAsString]; ok {
			reOrderedCmds = append(reOrderedCmds, oneCmd)
			continue
		}

		str := fmt.Sprintf("the command with the hash: %s could not be found", blkAsString)
		return nil, errors.New(str)
	}

	out, outErr := rep.cmdsBuilderFactory.Create().Create().WithMetaData(met).WithCommands(reOrderedCmds).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
