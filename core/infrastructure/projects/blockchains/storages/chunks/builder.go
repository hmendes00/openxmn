package chunks

import (
	"errors"
	"time"

	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/chunks"
	files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
)

type builder struct {
	ht        files.File
	chks      []files.File
	createdOn *time.Time
}

func createBuilder() chunk.ChunksBuilder {
	out := builder{
		ht:        nil,
		chks:      nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the ChunksBuilder instance
func (build *builder) Create() chunk.ChunksBuilder {
	build.ht = nil
	build.chks = nil
	build.createdOn = nil
	return build
}

// WithHashTree adds an HashTree file to the ChunksBuilder instance
func (build *builder) WithHashTree(ht files.File) chunk.ChunksBuilder {
	build.ht = ht
	return build
}

// WithChunks adds Chunks files to the ChunksBuilder instance
func (build *builder) WithChunks(fil []files.File) chunk.ChunksBuilder {
	build.chks = fil
	return build
}

// CreatedOn adds creation time to the ChunksBuilder instance
func (build *builder) CreatedOn(ts time.Time) chunk.ChunksBuilder {
	build.createdOn = &ts
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

	if build.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Chunks instance")
	}

	if len(build.chks) <= 0 {
		return nil, errors.New("there must be at least 1 file in the chunks in order to build a Chunks instance")
	}

	chks := []*concrete_files.File{}
	for _, oneChks := range build.chks {
		chks = append(chks, oneChks.(*concrete_files.File))
	}

	out := createChunks(build.ht.(*concrete_files.File), chks, *build.createdOn)
	return out, nil
}
