package infrastructure

import (
	"errors"
	"time"

	chunk "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	concrete_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
)

type chunksBuilder struct {
	ht        files.File
	chks      []files.File
	createdOn *time.Time
}

func createChunksBuilder() chunk.ChunksBuilder {
	out := chunksBuilder{
		ht:        nil,
		chks:      nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the ChunksBuilder instance
func (build *chunksBuilder) Create() chunk.ChunksBuilder {
	build.ht = nil
	build.chks = nil
	build.createdOn = nil
	return build
}

// WithHashTree adds an HashTree file to the ChunksBuilder instance
func (build *chunksBuilder) WithHashTree(ht files.File) chunk.ChunksBuilder {
	build.ht = ht
	return build
}

// WithChunks adds Chunks files to the ChunksBuilder instance
func (build *chunksBuilder) WithChunks(fil []files.File) chunk.ChunksBuilder {
	build.chks = fil
	return build
}

// CreatedOn adds creation time to the ChunksBuilder instance
func (build *chunksBuilder) CreatedOn(ts time.Time) chunk.ChunksBuilder {
	build.createdOn = &ts
	return build
}

// Now builds a new Chunks instance
func (build *chunksBuilder) Now() (chunk.Chunks, error) {
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
