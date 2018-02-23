package canada

import uuid "github.com/satori/go.uuid"

// Save represents a save canadian trust transaction
type Save interface {
	GetID() *uuid.UUID
	GetOrganizationID() *uuid.UUID
	GetInstitutionNumber() int
	GetSwiftCode() string
	GetDepositFees() float64
	GetName() string
	GetDescription() string
}
