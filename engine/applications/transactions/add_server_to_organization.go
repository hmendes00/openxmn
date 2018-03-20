package transactions

import uuid "github.com/satori/go.uuid"

// AddServerToOrganization represents transaction to add a server to an organization
type AddServerToOrganization struct {
	ServerID *uuid.UUID `json:"server_id"`
}
