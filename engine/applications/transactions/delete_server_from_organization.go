package transactions

import uuid "github.com/satori/go.uuid"

// DeleteServerFromOrganization represents transaction to delete a server from an organization
type DeleteServerFromOrganization struct {
	ServerID *uuid.UUID `json:"server_id"`
}
