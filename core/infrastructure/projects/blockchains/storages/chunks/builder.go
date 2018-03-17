package chunks

import (
	"errors"

	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/chunks"
	files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
)

type builder struct {
	ht   files.File
	chks []files.File
}

func createBuilder() chunk.Builder {
	out := builder{
		ht:   nil,
		chks: nil,
	}

	return &out
}

// Create initializes the ChunksBuilder instance
func (build *builder) Create() chunk.Builder {
	build.ht = nil
	build.chks = nil
	return build
}

// WithHashTree adds an HashTree file to the ChunksBuilder instance
func (build *builder) WithHashTree(ht files.File) chunk.Builder {
	build.ht = ht
	return build
}

// WithChunks adds Chunks files to the ChunksBuilder instance
func (build *builder) WithChunks(fil []files.File) chunk.Builder {
	build.chks = fil
	return build
}

// Now builds a new Chunks instance
func (build *builder) Now() (chunk.Chunks, error) {
	if build.ht == nil {
		return nil, errors.New("the hashtree file is mandatory in order to build a Chunks instance")
	}

	if build.chks == nil {
		return nil, errors.New("the chunks files are mandatory in order to build a Chunks instance")
	}

	if len(build.chks) <= 0 {
		return nil, errors.New("there must be at least 1 file in the chunks in order to build a Chunks instance")
	}

	chks := []*concrete_files.File{}
	for _, oneChks := range build.chks {
		chks = append(chks, oneChks.(*concrete_files.File))
	}

	out := createChunks(build.ht.(*concrete_files.File), chks)
	return out, nil
}
