package chunks

import (
	"time"

	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/chunks"
	files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
)

// Chunks represents a concrete Chunks implementation
type Chunks struct {
	HT   *concrete_files.File   `json:"hashtree"`
	Chks []*concrete_files.File `json:"chunks"`
	CrOn time.Time              `json:"created_on"`
}

func createChunks(ht *concrete_files.File, chks []*concrete_files.File, createdOn time.Time) chunk.Chunks {
	out := Chunks{
		HT:   ht,
		Chks: chks,
		CrOn: createdOn,
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

// CreatedOn returns the creation time
func (chks *Chunks) CreatedOn() time.Time {
	return chks.CrOn
}
