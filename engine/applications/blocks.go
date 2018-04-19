package main

import (
	"log"
	"time"

	agents "github.com/XMNBlockchain/openxmn/engine/applications/agents"
	apis "github.com/XMNBlockchain/openxmn/engine/applications/apis"
	databases "github.com/XMNBlockchain/openxmn/engine/applications/databases"
	processors_servers "github.com/XMNBlockchain/openxmn/engine/applications/processors/servers"
	processors_wealth "github.com/XMNBlockchain/openxmn/engine/applications/processors/wealth"
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	validated_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	processors "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/processors"
	servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_stored_commands "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/commands"
	concrete_stored_chunks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/chunks"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
	concrete_blocks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/blocks"
	concrete_blocks_validated "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/blocks/validated"
	concrete_commands "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/commands"
	concrete_bills "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/commands/bills"
	concrete_blockchain_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_processors "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/processors"
	concrete_chunks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/chunks"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_organizations "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/organizations"
	concrete_organizations_server "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/organizations/servers"
	concrete_stakes "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/organizations/stakes"
	concrete_safes "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/safes"
	concrete_servers "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/servers"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	concrete_wallets "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users/wallets"
	"github.com/gorilla/mux"
)

// Blocks represents a blocks application
type Blocks struct {
	api                   *apis.Blocks
	agent                 *agents.ValidateBlock
	chainedCmdProcessor   processors.ChainedCommands
	newSignedValidatedBlk <-chan validated_blocks.SignedBlock
}

// CreateBlocks creates a new Blocks instance
func CreateBlocks(
	router *mux.Router,
	routePrefix string,
	pk cryptography.PrivateKey,
	user users.User,
	databaseDirPath string,
	chkSizeInBytes int,
	chkFileExtension string,
	usersStake map[string]float64,
	signedBlkBufferSize int,
	validatedBlkBufferSize int,
	neededStakerPerBlk float64,
	waitBeforeRemovalTs time.Duration,
) *Blocks {

	//channels:
	newSignedBlock := make(chan blocks.SignedBlock, signedBlkBufferSize)
	newSignedValidatedBlk := make(chan validated_blocks.SignedBlock, validatedBlkBufferSize)

	//users related builder factories:
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	publicKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, htBuilderFactory, metaDataBuilderFactory)
	userSigsBuilderFactory := concrete_users.CreateSignaturesBuilderFactory(metaDataBuilderFactory)

	//transactions and blocks factories:
	blockMetaDataBuilderFactory := concrete_blockchain_metadata.CreateBuilderFactory()
	signedBlockBuilderFactory := concrete_blocks.CreateSignedBlockBuilderFactory(htBuilderFactory, blockMetaDataBuilderFactory)
	valBlkBuilderFactory := concrete_blocks_validated.CreateBlockBuilderFactory(htBuilderFactory, blockMetaDataBuilderFactory)
	signedValBlkBuilderFactory := concrete_blocks_validated.CreateSignedBlockBuilderFactory(htBuilderFactory, blockMetaDataBuilderFactory)

	//create the data structures:
	servs := map[string]servers.Server{}
	wals := map[string]wallets.Wallet{}

	//create the databases:
	orgDB := databases.CreateOrganization()
	safeDB := databases.CreateSafe()
	stakeDB := databases.CreateStake()
	tokDB := databases.CreateToken()
	userDB := databases.CreateUser()
	walDB := databases.CreateWallet()
	serverDB := databases.CreateServer()

	//create the objects related factories:
	walletBuilderFactory := concrete_wallets.CreateWalletBuilderFactory(metaDataBuilderFactory)
	organizationBuilderFactory := concrete_organizations.CreateOrganizationBuilderFactory(metaDataBuilderFactory)
	tokenBuilderFactory := concrete_tokens.CreateTokenBuilderFactory(metaDataBuilderFactory)
	safeBuilderFactory := concrete_safes.CreateSafeBuilderFactory(metaDataBuilderFactory)
	userBuilderFactory := concrete_users.CreateUserBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	stakeBuilderFactory := concrete_stakes.CreateStakeBuilderFactory(metaDataBuilderFactory)
	serverBuilderFactory := concrete_servers.CreateServerBuilderFactory()
	priceBuilderFactory := concrete_organizations_server.CreatePriceBuilderFactory()
	orgServerBuilderFactory := concrete_organizations_server.CreateServerBuilderFactory(metaDataBuilderFactory)

	//create the token accepted by our organization:
	tok, _ := tokenBuilderFactory.Create().Create().Now()

	//create the commands related factories:
	updateBuilderFactory := concrete_commands.CreateUpdateBuilderFactory()
	deleteBuilderFactory := concrete_commands.CreateDeleteBuilderFactory()
	insertBuilderFactory := concrete_commands.CreateInsertBuilderFactory()
	cmdBuilderFactory := concrete_commands.CreateCommandBuilderFactory()
	cmdsBuilderFactory := concrete_commands.CreateBuilderFactory()
	chainedCmdsBuilderFactory := concrete_commands.CreateChainedCommandsBuilderFactory(blockMetaDataBuilderFactory, htBuilderFactory)

	//create the command service:
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	fileService := concrete_files.CreateFileService(storedFileBuilderFactory)
	metaDataService := concrete_blockchain_metadata.CreateService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	htService := concrete_hashtrees.CreateHashTreeService(fileService, fileBuilderFactory)
	storedChkBuilderFactory := concrete_stored_chunks.CreateBuilderFactory()
	chksService := concrete_chunks.CreateService(htService, fileService, fileBuilderFactory, storedChkBuilderFactory)
	chksBuilderFactory := concrete_chunks.CreateBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, chkFileExtension)
	storedCmdsBuilderFactory := concrete_stored_commands.CreateBuilderFactory()
	cmdsService := concrete_commands.CreateService(metaDataService, chksService, chksBuilderFactory, storedCmdsBuilderFactory)

	//create the chained commands service:
	storedChainedCommandsBuilderFactory := concrete_stored_commands.CreateChainedCommandsBuilderFactory()
	chainedCommandsService := concrete_commands.CreateChainedCommandsService(metaDataService, cmdsService, fileService, fileBuilderFactory, storedChainedCommandsBuilderFactory)

	//create the bills related factories:
	payeeBuilderFactory := concrete_bills.CreatePayeeBuilderFactory()
	payerBuilderFactory := concrete_bills.CreatePayerBuilderFactory()
	billBuilderFactory := concrete_bills.CreateBillBuilderFactory(payeeBuilderFactory, payerBuilderFactory, servs, wals)

	//create the wealth processors:
	procCashSafe := processors_wealth.CreateCashSafe(safeDB, walDB, tokDB, walletBuilderFactory, updateBuilderFactory, deleteBuilderFactory, cmdBuilderFactory, cmdsBuilderFactory)
	procDeleteOrganization := processors_wealth.CreateDeleteOrganization(orgDB, cmdBuilderFactory, deleteBuilderFactory)
	procDeleteSafe := processors_wealth.CreateDeleteSafe(safeDB, cmdBuilderFactory, deleteBuilderFactory)
	procDeleteToken := processors_wealth.CreateDeleteToken(tokDB, cmdBuilderFactory, deleteBuilderFactory)
	procDeleteUser := processors_wealth.CreateDeleteUser(userDB, cmdBuilderFactory, deleteBuilderFactory)
	procInsertOrganization := processors_wealth.CreateInsertOrganization(orgDB, tokDB, organizationBuilderFactory, cmdBuilderFactory, insertBuilderFactory)
	procInsertToken := processors_wealth.CreateInsertToken(tokDB, userDB, safeDB, tokenBuilderFactory, safeBuilderFactory, cmdBuilderFactory, cmdsBuilderFactory, insertBuilderFactory, updateBuilderFactory)
	procInsertUser := processors_wealth.CreateInsertUser(userDB, userBuilderFactory, cmdBuilderFactory, insertBuilderFactory)
	procStakeTokenReversal := processors_wealth.CreateStakeTokenReversal(stakeDB, walDB, walletBuilderFactory, cmdBuilderFactory, cmdsBuilderFactory, updateBuilderFactory, deleteBuilderFactory)
	procStakeToken := processors_wealth.CreateStakeToken(orgDB, tokDB, walDB, stakeDB, stakeBuilderFactory, walletBuilderFactory, cmdBuilderFactory, cmdsBuilderFactory, insertBuilderFactory, updateBuilderFactory)
	procUpdateOrganization := processors_wealth.CreateUpdateOrganization(orgDB, tokDB, organizationBuilderFactory, cmdBuilderFactory, updateBuilderFactory)
	procUpdateSafe := processors_wealth.CreateUpdateSafe(safeDB, safeBuilderFactory, cmdBuilderFactory, updateBuilderFactory)
	procUpdateToken := processors_wealth.CreateUpdateToken(tokDB, tokenBuilderFactory, cmdBuilderFactory, updateBuilderFactory)
	procUpdateUser := processors_wealth.CreateUpdateUser(userDB, userBuilderFactory, cmdBuilderFactory, updateBuilderFactory)

	//create the servers processors:
	procSaveServer := processors_servers.CreateSaveServer(serverDB, orgDB, priceBuilderFactory, serverBuilderFactory, orgServerBuilderFactory, cmdBuilderFactory, insertBuilderFactory)
	procDelServer := processors_servers.CreateDeleteServer(serverDB, cmdBuilderFactory, deleteBuilderFactory)

	//create the transactions processor:
	trsProcessor := concrete_processors.CreateTransactions([]processors.Transaction{
		procCashSafe,
		procDeleteOrganization,
		procDeleteSafe,
		procDeleteToken,
		procDeleteUser,
		procInsertOrganization,
		procInsertToken,
		procInsertUser,
		procStakeTokenReversal,
		procStakeToken,
		procUpdateOrganization,
		procUpdateSafe,
		procUpdateToken,
		procUpdateUser,
		procSaveServer,
		procDelServer,
	}, cmdsBuilderFactory)

	//create the block processor:
	blkProcessor := concrete_processors.CreateBlock(
		tok,
		trsProcessor,
		billBuilderFactory,
		cmdBuilderFactory,
		cmdsBuilderFactory,
	)

	//create the chained commands processor:
	chainedCmdProcessor := concrete_processors.CreateChainedCommands(
		databaseDirPath,
		chainedCmdsBuilderFactory,
		chainedCommandsService,
		cmdsService,
		blkProcessor,
	)

	//create the blocks API:
	blocksAPI := apis.CreateBlocks(
		routePrefix,
		router,
		signedBlockBuilderFactory,
		newSignedBlock,
	)

	//create the block agents:
	agent := agents.CreateValidateBlock(
		pk,
		user,
		userSigBuilderFactory,
		valBlkBuilderFactory,
		signedValBlkBuilderFactory,
		userSigsBuilderFactory,
		neededStakerPerBlk,
		usersStake,
		waitBeforeRemovalTs,
		newSignedBlock,
		newSignedValidatedBlk,
	)

	out := Blocks{
		api:                   blocksAPI,
		agent:                 agent,
		chainedCmdProcessor:   chainedCmdProcessor,
		newSignedValidatedBlk: newSignedValidatedBlk,
	}

	return &out
}

// Execute execute the blocks application
func (blks *Blocks) Execute() {

	//execute the agent:
	go blks.agent.Execute()

	for {
		select {
		case oneValidatedBlk := <-blks.newSignedValidatedBlk:
			//process the validated block:
			procErr := blks.chainedCmdProcessor.Process(oneValidatedBlk)
			if procErr != nil {
				log.Fatalf("there was an error while processing the validated block (ID: %s): %s", oneValidatedBlk.GetMetaData().GetID().String(), procErr.Error())
			}

			break
		}

	}

}
