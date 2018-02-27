package objects

import uuid "github.com/satori/go.uuid"

// Organization represents an organization
type Organization struct {
	ID          *uuid.UUID         `json:"id"`
	Parent      *Organization      `json:"parent_organization"`
	Sym         *Symbol            `json:"symbol"`
	Creator     *User              `json:"creator"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Wal         map[string]*Wallet `json:"wallet"`
}
