package infrastructure

import (
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
