package chunks

import (
	chunk "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/chunks"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// Chunks represents a concrete Chunks implementation
type Chunks struct {
	HT   *concrete_files.File   `json:"hashtree"`
	Chks []*concrete_files.File `json:"chunks"`
}

func createChunks(ht *concrete_files.File, chks []*concrete_files.File) chunk.Chunks {
	out := Chunks{
		HT:   ht,
		Chks: chks,
	}

	return &out
}

// GetHashTree returns the hashtree file
func (chks *Chunks) GetHashTree() files.File {
	return chks.HT
}

// GetChunks returns the Chunks files
func (chks *Chunks) GetChunks() []files.File {
	out := []files.File{}
	for _, oneChk := range chks.Chks {
		out = append(out, oneChk)
	}

	return out
}
