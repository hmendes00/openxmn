package infrastructure

import (
	"time"

	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
	uuid "github.com/satori/go.uuid"
)

type object struct {
	id   *uuid.UUID
	crOn time.Time
	sig  stored_files.File
	chks stored_chunks.Chunks
}

func createObject(id *uuid.UUID, chks stored_chunks.Chunks, crOn time.Time) objects.Object {
	out := object{
		id:   id,
		crOn: crOn,
		sig:  nil,
		chks: chks,
	}

	return &out
}

func createObjectWithSignature(id *uuid.UUID, chks stored_chunks.Chunks, sig stored_files.File, crOn time.Time) objects.Object {
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

// GetChunks returns the chunks file, if any
func (obj *object) GetChunks() stored_chunks.Chunks {
	return obj.chks
}

// HasSignature returns true if there is a signature, false otherwise
func (obj *object) HasSignature() bool {
	return obj.sig != nil
}

// GetSignature returns the signature file, if any
func (obj *object) GetSignature() stored_files.File {
	return obj.sig
}
