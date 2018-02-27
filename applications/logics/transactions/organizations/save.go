package organizations

import uuid "github.com/satori/go.uuid"

// Save represents a save organization transaction
type Save struct {
	ID          *uuid.UUID `json:"id"`
	TokenID     *uuid.UUID `json:"token_id"`
	UserID      *uuid.UUID `json:"user_id"`
	ParentID    *uuid.UUID `json:"parent_organization_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}

// CreateSave creates a new Save instance
func CreateSave(id *uuid.UUID, tokenID *uuid.UUID, userID *uuid.UUID, name string, description string) *Save {
	out := Save{
		ID:          id,
		TokenID:     tokenID,
		UserID:      userID,
		ParentID:    nil,
		Name:        name,
		Description: description,
	}

	return &out
}

// CreateSaveWithOrganizationID creates a new Save instance with a parent organization
func CreateSaveWithOrganizationID(id *uuid.UUID, tokenID *uuid.UUID, userID *uuid.UUID, parentID *uuid.UUID, name string, description string) *Save {
	out := Save{
		ID:          id,
		TokenID:     tokenID,
		UserID:      userID,
		ParentID:    parentID,
		Name:        name,
		Description: description,
	}

	return &out
}

// GetID returns the ID
func (sav *Save) GetID() *uuid.UUID {
	return sav.ID
}

// GetTokenID returns the token ID
func (sav *Save) GetTokenID() *uuid.UUID {
	return sav.TokenID
}

// GetUserID returns the user ID
func (sav *Save) GetUserID() *uuid.UUID {
	return sav.UserID
}

// HasParentID returns true if there is a parent organization ID, false otherwise
func (sav *Save) HasParentID() bool {
	return sav.ParentID != nil
}

// GetParentID returns the parent organization ID, if any
func (sav *Save) GetParentID() *uuid.UUID {
	return sav.ParentID
}

// GetName returns the name of the organization
func (sav *Save) GetName() string {
	return sav.Name
}

// GetDescription returns the description of the organization
func (sav *Save) GetDescription() string {
	return sav.Description
}
