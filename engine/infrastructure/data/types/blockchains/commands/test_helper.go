package commands

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	concrete_stored_commands "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/commands"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_chunks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/chunks"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	uuid "github.com/satori/go.uuid"
)

// SomeDataForTests represents some data, only for tests
type SomeDataForTests struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateChainedCommandsServiceForTests creates a new ChainedCommandsService instance for tests
func CreateChainedCommandsServiceForTests() commands.ChainedCommandsService {
	metaDataService := concrete_metadata.CreateServiceForTests()
	cmdsService := CreateServiceForTests()
	fileService := concrete_files.CreateFileServiceForTests()
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	storedChainedCommandsBuilderFactory := concrete_stored_commands.CreateChainedCommandsBuilderFactory()
	out := CreateChainedCommandsService(metaDataService, cmdsService, fileService, fileBuilderFactory, storedChainedCommandsBuilderFactory)
	return out
}

// CreateChainedCommandsRepositoryForTests creates a new ChainedCommandsRepository instance for tests
func CreateChainedCommandsRepositoryForTests() commands.ChainedCommandsRepository {
	metaDataRepository := concrete_metadata.CreateRepositoryForTests()
	cmdsRepository := CreateRepositoryForTests()
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	cmdsBuilderFactory := CreateChainedCommandsBuilderFactoryForTests()
	out := CreateChainedCommandsRepository(metaDataRepository, cmdsRepository, fileRepository, cmdsBuilderFactory)
	return out
}

// CreateRepositoryForTests creates a new Repository instance for tests
func CreateRepositoryForTests() commands.Repository {
	metaDataRepository := concrete_metadata.CreateRepositoryForTests()
	chunksRepository := concrete_chunks.CreateRepositoryForTests()
	cmdsBuilderFactory := CreateBuilderFactoryForTests()
	out := CreateRepository(metaDataRepository, chunksRepository, cmdsBuilderFactory)
	return out
}

// CreateServiceForTests creates a new Service instance for tests
func CreateServiceForTests() commands.Service {
	metaDataService := concrete_metadata.CreateServiceForTests()
	chksService := concrete_chunks.CreateServiceForTests()
	chksBuilderFactory := concrete_chunks.CreateBuilderFactoryForTests()
	storedCmdsBuilderFactory := concrete_stored_commands.CreateBuilderFactoryForTests()
	out := CreateService(metaDataService, chksService, chksBuilderFactory, storedCmdsBuilderFactory)
	return out
}

// CreateBuilderFactoryForTests creates a new BuilderFactory for tests
func CreateBuilderFactoryForTests() commands.BuilderFactory {
	out := CreateBuilderFactory()
	return out
}

// CreateChainedCommandsBuilderFactoryForTests creates a new ChainedCommandsBuilderFactory instance for tests
func CreateChainedCommandsBuilderFactoryForTests() commands.ChainedCommandsBuilderFactory {
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactoryForTests()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	out := CreateChainedCommandsBuilderFactory(metaDataBuilderFactory, htBuilderFactory)
	return out
}

// CreateChainedCommandsForTests creates a new ChainedCommands instance for tests
func CreateChainedCommandsForTests() *ChainedCommands {
	metaData := concrete_metadata.CreateMetaDataForTests()
	cmds := CreateCommandsForTests()
	prevID := uuid.NewV4()
	rootID := uuid.NewV4()

	out := createChainedCommands(metaData, cmds, &prevID, &rootID)
	return out.(*ChainedCommands)
}

// CreateCommandsForTests creates a new Commands instance for tests
func CreateCommandsForTests() *Commands {
	firstCmd := CreateCommandForTests()
	secondCmd := CreateCommandForTests()
	thirdCmd := CreateCommandForTests()

	id := uuid.NewV4()
	ts := time.Now().UTC()
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks([][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(ts.UnixNano()))),
		firstCmd.GetMetaData().GetHashTree().GetHash().Get(),
		secondCmd.GetMetaData().GetHashTree().GetHash().Get(),
		thirdCmd.GetMetaData().GetHashTree().GetHash().Get(),
	}).Now()

	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).CreatedOn(ts).WithHashTree(ht).Now()

	out := createCommands(met.(*concrete_metadata.MetaData), []*Command{
		firstCmd,
		secondCmd,
		thirdCmd,
	})

	return out.(*Commands)
}

// CreateCommandForTests creates a new Command instance for tests
func CreateCommandForTests() *Command {
	ins := CreateInsertForTests()

	id := uuid.NewV4()
	ts := time.Now().UTC()
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks([][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(ts.UnixNano()))),
		ins.GetJS(),
	}).Now()

	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).CreatedOn(ts).WithHashTree(ht).Now()

	out := createCommandWithInsert(met.(*concrete_metadata.MetaData), ins)
	return out.(*Command)
}

// CreateInsertForTests creates a new Insert instance for tests
func CreateInsertForTests() *Insert {
	ins := SomeDataForTests{
		Name:        "Roger",
		Description: "Cyr",
	}

	js, _ := json.Marshal(ins)
	out := createInsert(js)
	return out.(*Insert)
}

// CreateUpdateForTests creates a new Update instance for tests
func CreateUpdateForTests() *Update {
	old := SomeDataForTests{
		Name:        "Roger",
		Description: "Cyr",
	}

	new := SomeDataForTests{
		Name:        "Steve",
		Description: "Rodrigue",
	}

	oldJS, _ := json.Marshal(old)
	newJS, _ := json.Marshal(new)

	out := createUpdate(oldJS, newJS)
	return out.(*Update)
}

// CreateDeleteForTests creates a new Delete instance for tests
func CreateDeleteForTests() *Delete {
	ins := SomeDataForTests{
		Name:        "Roger",
		Description: "Cyr",
	}

	js, _ := json.Marshal(ins)
	out := createDelete(js)
	return out.(*Delete)
}

// CreateErrorForTests creates a new Error instance for tests
func CreateErrorForTests() *Error {
	trsID := uuid.NewV4()
	code := rand.Int() % 20
	message := "this is an error!"
	out := createError(&trsID, code, message)
	return out.(*Error)
}
