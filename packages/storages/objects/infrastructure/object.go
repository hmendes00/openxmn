package infrastructure

import (
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type object struct {
	metaData stored_files.File
	chks     stored_chunks.Chunks
}

func createObject(metaData stored_files.File) objs.Object {
	out := object{
		metaData: metaData,
		chks:     nil,
	}

	return &out
}

func createObjectWithChunks(metaData stored_files.File, chks stored_chunks.Chunks) objs.Object {
	out := object{
		metaData: metaData,
		chks:     chks,
	}

	return &out
}

// GetMetaData returns the MetaData
func (obj *object) GetMetaData() stored_files.File {
	return obj.metaData
}

// HasChunks returns true if there is chunks, false otherwise
func (obj *object) HasChunks() bool {
	return obj.chks != nil
}

// GetChunks returns the chunks file, if any
func (obj *object) GetChunks() stored_chunks.Chunks {
	return obj.chks
}
