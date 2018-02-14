package infrastructure

import (
	"time"

	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents a concrete metadata implementation
type MetaData struct {
	ID   *uuid.UUID                `json:"id"`
	Sig  *concrete_users.Signature `json:"signature"`
	CrOn time.Time                 `json:"created_on"`
}

func createMetaData(id *uuid.UUID, crOn time.Time) objs.MetaData {
	out := MetaData{
		ID:   id,
		Sig:  nil,
		CrOn: crOn,
	}

	return &out
}

func createMetaDataWithSignature(id *uuid.UUID, sig *concrete_users.Signature, crOn time.Time) objs.MetaData {
	out := MetaData{
		ID:   id,
		Sig:  sig,
		CrOn: crOn,
	}

	return &out
}

// GetID returns the ID
func (met *MetaData) GetID() *uuid.UUID {
	return met.ID
}

// HasSignature returns true if there is a signature, false otherwise
func (met *MetaData) HasSignature() bool {
	return met.Sig != nil
}

// GetSignature returns the signature, if any
func (met *MetaData) GetSignature() users.Signature {
	return met.Sig
}

// CreatedOn returns the creation time
func (met *MetaData) CreatedOn() time.Time {
	return met.CrOn
}
