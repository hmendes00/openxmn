package safes

import (
	"errors"
	"time"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	uuid "github.com/satori/go.uuid"
)

type safeBuilder struct {
	metaDataBuilderFactory metadata.BuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	ciph                   cryptography.Cipher
	crOn                   *time.Time
	lstUpOn                *time.Time
}

func createSafeBuilder(metaDataBuilderFactory metadata.BuilderFactory) safes.SafeBuilder {
	out := safeBuilder{
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:      nil,
		met:     nil,
		ciph:    nil,
		crOn:    nil,
		lstUpOn: nil,
	}

	return &out
}

// Create initializes the safe
func (build *safeBuilder) Create() safes.SafeBuilder {
	build.id = nil
	build.met = nil
	build.ciph = nil
	build.crOn = nil
	build.lstUpOn = nil
	return build
}

// WithID adds an ID to the SafeBuilder instance
func (build *safeBuilder) WithID(id *uuid.UUID) safes.SafeBuilder {
	build.id = id
	return build
}

// WithMetaData adds metadata to the SafeBuilder instance
func (build *safeBuilder) WithMetaData(met metadata.MetaData) safes.SafeBuilder {
	build.met = met
	return build
}

// WithCipher adds a cipher to the SafeBuilder instance
func (build *safeBuilder) WithCipher(cipher cryptography.Cipher) safes.SafeBuilder {
	build.ciph = cipher
	return build
}

// CreatedOn adds a creation time to the SafeBuilder instance
func (build *safeBuilder) CreatedOn(crOn time.Time) safes.SafeBuilder {
	build.crOn = &crOn
	return build
}

// LastUpdatedOn adds a update time to the SafeBuilder instance
func (build *safeBuilder) LastUpdatedOn(lstUpOn time.Time) safes.SafeBuilder {
	build.lstUpOn = &lstUpOn
	return build
}

// Now builds a new Safe instance
func (build *safeBuilder) Now() (safes.Safe, error) {

	if build.ciph == nil {
		return nil, errors.New("the cipher is mandatory in order to build a Safe instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a Safe instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a Safe instance")
		}

		metBuilder := build.metaDataBuilderFactory.Create().Create().WithID(build.id).CreatedOn(*build.crOn)
		if build.lstUpOn != nil {
			metBuilder.LastUpdatedOn(*build.lstUpOn)
		}

		met, metErr := metBuilder.Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	out := createSafe(build.met.(*concrete_metadata.MetaData), build.ciph.(*concrete_cryptography.Cipher))
	return out, nil
}
