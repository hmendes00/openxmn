package chunks

import (
	"bytes"
	"encoding/gob"

	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/chunks"
	files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/files"
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
)

type chunks struct {
	ht   hashtrees.HashTree
	chks []files.File
}

func createChunks(ht hashtrees.HashTree, chks []files.File) chunk.Chunks {
	out := chunks{
		ht:   ht,
		chks: chks,
	}

	return &out
}

// GetHashTree returns the HashTree
func (chks *chunks) GetHashTree() hashtrees.HashTree {
	return chks.ht
}

// GetChunks returns the file chunks
func (chks *chunks) GetChunks() []files.File {
	return chks.chks
}

// GetData returns the combined re-ordered data
func (chks *chunks) GetData() ([]byte, error) {
	//combine the chunks data:
	trsData := [][]byte{}
	for _, oneFileChk := range chks.chks {
		trsData = append(trsData, oneFileChk.GetData())
	}

	//re-order the data:
	reOrderedData, reOrderedDataErr := chks.GetHashTree().Order(trsData)
	if reOrderedDataErr != nil {
		return nil, reOrderedDataErr
	}

	//combine the data:
	matrixData := bytes.Join(reOrderedData, []byte(""))
	return matrixData, nil
}

// Marshal re-create the object bashed on the chunks and the hashtree
func (chks *chunks) Marshal(v interface{}) error {
	//re-create the object:
	matrixData, matrixDataErr := chks.GetData()
	if matrixDataErr != nil {
		return matrixDataErr
	}

	rdBuf := bytes.NewReader(matrixData)
	dec := gob.NewDecoder(rdBuf)
	decErr := dec.Decode(v)
	if decErr != nil {
		return decErr
	}

	return nil
}
