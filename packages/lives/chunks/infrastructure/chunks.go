package infrastructure

import (
	"bytes"
	"encoding/gob"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	chunk "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
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

// Marshal re-create the object bashed on the chunks and the hashtree
func (chks *chunks) Marshal(v interface{}) error {
	//combine the chunks data:
	trsData := [][]byte{}
	for _, oneFileChk := range chks.chks {
		trsData = append(trsData, oneFileChk.GetData())
	}

	//re-order the data:
	reOrderedData, reOrderedDataErr := chks.GetHashTree().Order(trsData)
	if reOrderedDataErr != nil {
		return reOrderedDataErr
	}

	//re-create the object:
	matrixData := bytes.Join(reOrderedData, []byte(""))
	rdBuf := bytes.NewReader(matrixData)
	dec := gob.NewDecoder(rdBuf)
	decErr := dec.Decode(v)
	if decErr != nil {
		return decErr
	}

	return nil
}
