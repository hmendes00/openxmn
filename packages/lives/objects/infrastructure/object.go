package infrastructure

import (
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

type object struct {
	met  objs.MetaData
	chks chunks.Chunks
}

func createObject(met objs.MetaData) objs.Object {
	out := object{
		met:  met,
		chks: nil,
	}

	return &out
}

func createObjectWithChunks(met objs.MetaData, chks chunks.Chunks) objs.Object {
	out := object{
		met:  met,
		chks: chks,
	}

	return &out
}

// GetMetaData returns the MetaData
func (obj *object) GetMetaData() objs.MetaData {
	return obj.met
}

// HasChunks returns true if there is chunks. false otherwise
func (obj *object) HasChunks() bool {
	return obj.chks != nil
}

// GetChunks returns the Chunks
func (obj *object) GetChunks() chunks.Chunks {
	return obj.chks
}
