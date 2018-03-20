package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

// SignaturesBuilder represents a SignaturesBulder
type SignaturesBuilder interface {
	Create() SignaturesBuilder
	WithID(id *uuid.UUID) SignaturesBuilder
	WithMetaData(met metadata.MetaData) SignaturesBuilder
	WithSignatures(sigs []Signature) SignaturesBuilder
	CreatedOn(crOn time.Time) SignaturesBuilder
	Now() (Signatures, error)
}
