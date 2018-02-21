package infrastructure

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/lives/metadata/domain"
	signed_trs "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
)

// AtomicTransactionRepository represents a concrete AtomicTransactionRepository implementation
type AtomicTransactionRepository struct {
	metaDataRepository      metadata.MetaDataRepository
	userSigRepository       users.SignatureRepository
	htRepository            hashtrees.HashTreeRepository
	trsRepository           transactions.TransactionRepository
	atomicTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory
}

// CreateAtomicTransactionRepository creates a new AtomicTransactionRepository instance
func CreateAtomicTransactionRepository(
	metaDataRepository metadata.MetaDataRepository,
	userSigRepository users.SignatureRepository,
	htRepository hashtrees.HashTreeRepository,
	trsRepository transactions.TransactionRepository,
	atomicTrsBuilderFactory signed_trs.AtomicTransactionBuilderFactory,
) signed_trs.AtomicTransactionRepository {
	out := AtomicTransactionRepository{
		metaDataRepository:      metaDataRepository,
		userSigRepository:       userSigRepository,
		htRepository:            htRepository,
		trsRepository:           trsRepository,
		atomicTrsBuilderFactory: atomicTrsBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a AtomicTransaction instance
func (rep *AtomicTransactionRepository) Retrieve(dirPath string) (signed_trs.AtomicTransaction, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the signature:
	sig, sigErr := rep.userSigRepository.Retrieve(dirPath)
	if sigErr != nil {
		return nil, sigErr
	}

	//retrieve the hashtree:
	ht, htErr := rep.htRepository.Retrieve(dirPath)
	if htErr != nil {
		return nil, htErr
	}

	//retrieve the transactions:
	trsDirPath := filepath.Join(dirPath, "transactions")
	trs, trsErr := rep.trsRepository.RetrieveAll(trsDirPath)
	if trsErr != nil {
		return nil, trsErr
	}

	trsMap := map[string]transactions.Transaction{}
	for _, oneTrs := range trs {
		trsIDAsString := hex.EncodeToString(oneTrs.GetID().Bytes())
		trsMap[trsIDAsString] = oneTrs
	}

	//re-order the transactions:
	trsIDs := [][]byte{}
	for _, oneTrs := range trs {
		trsIDs = append(trsIDs, oneTrs.GetID().Bytes())
	}

	orderedData, orderedErr := ht.Order(trsIDs)
	if orderedErr != nil {
		return nil, orderedErr
	}

	//create the new trs list based on the ordered data:
	orderedTrs := []transactions.Transaction{}
	for _, oneOrderedData := range orderedData {
		trsIDAsString := hex.EncodeToString(oneOrderedData)
		if oneTrs, ok := trsMap[trsIDAsString]; ok {
			orderedTrs = append(orderedTrs, oneTrs)
			continue
		}

		str := fmt.Sprintf("the ordered transaction ID (%s), in the HashTree, cannot be found.  Path: %s. \n\n%v", trsIDAsString, dirPath, trsMap)
		return nil, errors.New(str)
	}

	//build the atomic transaction:
	id := met.GetID()
	createdOn := met.CreatedOn()
	atomicTrs, atomicTrsErr := rep.atomicTrsBuilderFactory.Create().Create().WithID(id).CreatedOn(createdOn).WithSignature(sig).WithTransactions(orderedTrs).Now()
	if atomicTrsErr != nil {
		return nil, atomicTrsErr
	}

	return atomicTrs, nil
}

// RetrieveAll retrieves a []AtomicTransaction instances
func (rep *AtomicTransactionRepository) RetrieveAll(dirPath string) ([]signed_trs.AtomicTransaction, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	out := []signed_trs.AtomicTransaction{}
	for _, oneFile := range files {
		if !oneFile.IsDir() {
			continue
		}

		oneDirPath := filepath.Join(dirPath, oneFile.Name())
		oneTrs, oneTrsErr := rep.Retrieve(oneDirPath)
		if oneTrsErr != nil {
			return nil, oneTrsErr
		}

		out = append(out, oneTrs)
	}

	return out, nil
}
