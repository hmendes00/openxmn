package infrastructure

import (
	"time"

	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
	uuid "github.com/satori/go.uuid"
)

type object struct {
	ht   stored_files.File
	id   *uuid.UUID
	crOn time.Time
	sig  stored_files.File
	chks stored_chunks.Chunks
}

func createObject(ht stored_files.File, id *uuid.UUID, crOn time.Time) objects.Object {
	out := object{
		ht:   ht,
		id:   id,
		crOn: crOn,
		sig:  nil,
		chks: nil,
	}

	return &out
}

func createObjectWithSignature(ht stored_files.File, id *uuid.UUID, crOn time.Time, sig stored_files.File) objects.Object {
	out := object{
		ht:   ht,
		id:   id,
		crOn: crOn,
		sig:  sig,
		chks: nil,
	}

	return &out
}

func createObjectWithChunks(ht stored_files.File, id *uuid.UUID, crOn time.Time, chks stored_chunks.Chunks) objects.Object {
	out := object{
		ht:   ht,
		id:   id,
		crOn: crOn,
		sig:  nil,
		chks: chks,
	}

	return &out
}

func createObjectWithSignatureWithChunks(ht stored_files.File, id *uuid.UUID, crOn time.Time, sig stored_files.File, chks stored_chunks.Chunks) objects.Object {
	out := object{
		ht:   ht,
		id:   id,
		crOn: crOn,
		sig:  sig,
		chks: chks,
	}

	return &out
}

// GetHashTree returns the hashtree file
func (obj *object) GetHashTree() stored_files.File {
	return obj.ht
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
