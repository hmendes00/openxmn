package filesservers

import uuid "github.com/satori/go.uuid"

// StoreFileOnServer represents a transaction used to store a file on server
type StoreFileOnServer struct {
	ID       *uuid.UUID `json:"id"`
	ServerID *uuid.UUID `json:"server_id"`
	FilePath string     `json:"file_path"`
}
