package files

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	uuid "github.com/satori/go.uuid"
)

// File represents a concrete file implementation
type File struct {
	ID   *uuid.UUID `json:"id"`
	Typ  string     `json:"type"`
	Hash string     `json:"hash"`
}

func createFile(id *uuid.UUID, typ string, hash string) files.File {
	out := File{
		ID:   id,
		Typ:  typ,
		Hash: hash,
	}

	return &out
}

// GetID returns the ID
func (fil *File) GetID() *uuid.UUID {
	return fil.ID
}

// GetType returns the type
func (fil *File) GetType() string {
	return fil.Typ
}

// GetHash returns the hash
func (fil *File) GetHash() string {
	return fil.Hash
}
