package remote

import (
	"hash"

	servers "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/servers"
)

// File represents a remote file
type File struct {
	Path    string           `json:"path"`
	Hash    hash.Hash        `json:"hash"`
	Servers *servers.Servers `json:"servers"`
}

// CreateFile returns a file instance
func CreateFile(path string, hash hash.Hash, serv *servers.Servers) *File {
	out := File{
		Path:    path,
		Hash:    hash,
		Servers: serv,
	}

	return &out
}

// GetPath returns the file path
func (fil *File) GetPath() string {
	return fil.Path
}

// GetHash returns the file hash
func (fil *File) GetHash() hash.Hash {
	return fil.Hash
}

// GetServers returns the servers that host the file
func (fil *File) GetServers() *servers.Servers {
	return fil.Servers
}
