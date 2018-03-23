package users

import (
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SignaturesRepository represents a concrete SignaturesRepository implementation
type SignaturesRepository struct {
	metaDataRepository metadata.MetaDataRepository
	sigsRepository     users.SignatureRepository
	sigsBuilderFactory users.SignaturesBuilderFactory
}

// CreateSignaturesRepository creates a new SignaturesRepository instance
func CreateSignaturesRepository(metaDataRepository metadata.MetaDataRepository, sigsRepository users.SignatureRepository, sigsBuilderFactory users.SignaturesBuilderFactory) users.SignaturesRepository {
	out := SignaturesRepository{
		metaDataRepository: metaDataRepository,
		sigsRepository:     sigsRepository,
		sigsBuilderFactory: sigsBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a Signatures instance
func (rep *SignaturesRepository) Retrieve(dirPath string) (users.Signatures, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the signatures:
	sigsPath := filepath.Join(dirPath, "signatures")
	sigs, sigsErr := rep.sigsRepository.RetrieveAll(sigsPath)
	if sigsErr != nil {
		return nil, sigsErr
	}

	//retrieve the blocks:
	blocks := [][]byte{
		met.GetID().Bytes(),
		[]byte(strconv.Itoa(int(met.CreatedOn().UnixNano()))),
	}

	sigsMap := map[string]users.Signature{}
	for _, oneSig := range sigs {
		hash := oneSig.GetMetaData().GetHashTree().GetHash()
		sigsMap[hash.String()] = oneSig
		blocks = append(blocks, hash.Get())
	}

	//re-order the blocks:
	reOrderedBlks, reOrderedBlksErr := met.GetHashTree().Order(blocks)
	if reOrderedBlksErr != nil {
		return nil, reOrderedBlksErr
	}

	//re-order the signatures:
	reOrderedSigs := []users.Signature{}
	for _, oneBlk := range reOrderedBlks[2:] {
		blkAsString := hex.EncodeToString(oneBlk)
		if oneSig, ok := sigsMap[blkAsString]; ok {
			reOrderedSigs = append(reOrderedSigs, oneSig)
			continue
		}

		str := fmt.Sprintf("the user signature with the hash: %s could not be found", blkAsString)
		return nil, errors.New(str)
	}

	//build the user signatures:
	userSigs, userSigsErr := rep.sigsBuilderFactory.Create().Create().WithMetaData(met).WithSignatures(reOrderedSigs).Now()
	if userSigsErr != nil {
		return nil, userSigsErr
	}

	return userSigs, nil
}
