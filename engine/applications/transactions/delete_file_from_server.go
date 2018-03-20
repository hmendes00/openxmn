package transactions

import uuid "github.com/satori/go.uuid"

// DeleteFileFromServer represents a transaction used to delete a file from a server
type DeleteFileFromServer struct {
	StorageID *uuid.UUID `json:"storage_id"`
}
