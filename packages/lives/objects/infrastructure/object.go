package infrastructure

import (
	"time"

	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

type object struct {
	id   *uuid.UUID
	crOn time.Time
	sig  users.Signature
	chks chunks.Chunks
}

func createObject(id *uuid.UUID, createdOn time.Time) objs.Object {
	out := object{
		id:   id,
		crOn: createdOn,
		sig:  nil,
		chks: nil,
	}

	return &out
}

func createObjectWithChunks(id *uuid.UUID, createdOn time.Time, chks chunks.Chunks) objs.Object {
	out := object{
		id:   id,
		crOn: createdOn,
		sig:  nil,
		chks: chks,
	}

	return &out
}

func createObjectWithChunksWithSignature(id *uuid.UUID, createdOn time.Time, chks chunks.Chunks, sig users.Signature) objs.Object {
	out := object{
		id:   id,
		crOn: createdOn,
		sig:  sig,
		chks: chks,
	}

	return &out
}

func createObjectWithSignature(id *uuid.UUID, createdOn time.Time, sig users.Signature) objs.Object {
	out := object{
		id:   id,
		crOn: createdOn,
		sig:  sig,
		chks: nil,
	}

	return &out
}

// GetID returns the ID
func (obj *object) GetID() *uuid.UUID {
	return obj.id
}

// CreatedOn returns the creation ts:
func (obj *object) CreatedOn() time.Time {
	return obj.crOn
}

// HasSignature returns true if there is a signature, false otherwise
func (obj *object) HasSignature() bool {
	return obj.sig != nil
}

// GetSignature returns the signature, if any
func (obj *object) GetSignature() users.Signature {
	return obj.sig
}

// HasChunks returns true if there is chunks. false otherwise
func (obj *object) HasChunks() bool {
	return obj.chks != nil
}

// GetChunks returns the Chunks
func (obj *object) GetChunks() chunks.Chunks {
	return obj.chks
}
