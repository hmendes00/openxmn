package infrastructure

import (
	"time"

	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
	uuid "github.com/satori/go.uuid"
)

type object struct {
	id   *uuid.UUID
	crOn time.Time
	sig  stored_files.File
	chks stored_chunks.Chunks
}

func createObject(id *uuid.UUID, crOn time.Time) objs.Object {
	out := object{
		id:   id,
		crOn: crOn,
		sig:  nil,
		chks: nil,
	}

	return &out
}

func createObjectWithSignature(id *uuid.UUID, crOn time.Time, sig stored_files.File) objs.Object {
	out := object{
		id:   id,
		crOn: crOn,
		sig:  sig,
		chks: nil,
	}

	return &out
}

func createObjectWithChunks(id *uuid.UUID, crOn time.Time, chks stored_chunks.Chunks) objs.Object {
	out := object{
		id:   id,
		crOn: crOn,
		sig:  nil,
		chks: chks,
	}

	return &out
}

func createObjectWithSignatureWithChunks(id *uuid.UUID, crOn time.Time, sig stored_files.File, chks stored_chunks.Chunks) objs.Object {
	out := object{
		id:   id,
		crOn: crOn,
		sig:  sig,
		chks: chks,
	}

	return &out
}

// GetID returns the ID
func (obj *object) GetID() *uuid.UUID {
	return obj.id
}

// CreatedOn returns the creation time
func (obj *object) CreatedOn() time.Time {
	return obj.crOn
}

// HasSignature returns true if there is a signature, false otherwise
func (obj *object) HasSignature() bool {
	return obj.sig != nil
}

// GetSignature returns the signature file, if any
func (obj *object) GetSignature() stored_files.File {
	return obj.sig
}

// HasChunks returns true if there is chunks, false otherwise
func (obj *object) HasChunks() bool {
	return obj.chks != nil
}

// GetChunks returns the chunks file, if any
func (obj *object) GetChunks() stored_chunks.Chunks {
	return obj.chks
}
