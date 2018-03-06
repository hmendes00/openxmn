package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// SignaturesBuilder represents a SignaturesBulder
type SignaturesBuilder interface {
	Create() SignaturesBuilder
	WithID(id *uuid.UUID) SignaturesBuilder
	WithSignatures(sig []Signature) SignaturesBuilder
	CreatedOn(crOn time.Time) SignaturesBuilder
	Now() (Signatures, error)
}
