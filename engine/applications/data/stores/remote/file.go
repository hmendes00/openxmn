package remote

import (
	fil "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// File represents a stored remote file
type File struct {
	Servs []*fil.File `json:"servers"`
	Fil   *fil.File   `json:"file"`
}

// CreateFile creates a new file instance
func CreateFile(servers []*fil.File, file *fil.File) *File {
	out := File{
		Servs: servers,
		Fil:   file,
	}

	return &out
}

// GetServers returns the files that contains the servers that contains the file
func (fil *File) GetServers() []*fil.File {
	return fil.Servs
}

// GetFile returns the file
func (fil *File) GetFile() *fil.File {
	return fil.Fil
}
