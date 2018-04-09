package users

import (
	"errors"
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	user "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	uuid "github.com/satori/go.uuid"
)

type signaturesBuilder struct {
	metaDataBuilderFactory metadata.BuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	sigs                   []user.Signature
	crOn                   *time.Time
}

func createSignaturesBuilder(metaDataBuilderFactory metadata.BuilderFactory) user.SignaturesBuilder {
	out := signaturesBuilder{
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:   nil,
		met:  nil,
		sigs: nil,
		crOn: nil,
	}

	return &out
}

// Create initializes the SignaturesBuilder
func (build *signaturesBuilder) Create() user.SignaturesBuilder {
	build.id = nil
	build.met = nil
	build.sigs = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the SignaturesBuilder
func (build *signaturesBuilder) WithID(id *uuid.UUID) user.SignaturesBuilder {
	build.id = id
	return build
}

// WithMetaData adds a MetaData to the SignaturesBuilder
func (build *signaturesBuilder) WithMetaData(met metadata.MetaData) user.SignaturesBuilder {
	build.met = met
	return build
}

// WithSignatures adds []Signature to the SignaturesBuilder
func (build *signaturesBuilder) WithSignatures(sigs []user.Signature) user.SignaturesBuilder {
	build.sigs = sigs
	return build
}

// CreatedOn adds creation time to the SignaturesBuilder instance
func (build *signaturesBuilder) CreatedOn(crOn time.Time) user.SignaturesBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a new Signatures instance
func (build *signaturesBuilder) Now() (user.Signatures, error) {
	if build.sigs == nil {
		return nil, errors.New("the []Signature are mandatory in order to build a Signatures instance")
	}

	if len(build.sigs) <= 0 {
		return nil, errors.New("there must be at least 1 Signature instance in order to build a Signatures instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a Signatures instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a Signatures instance")
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).CreatedOn(*build.crOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Signatures instance")
	}

	sigs := []*Signature{}
	for _, oneSig := range build.sigs {
		sigs = append(sigs, oneSig.(*Signature))
	}

	out := createSignatures(build.met.(*concrete_metadata.MetaData), sigs)
	return out, nil
}
