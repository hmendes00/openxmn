package tokens

import (
	"errors"
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

type tokenBuilder struct {
	metaDataBuilderFactory metadata.BuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	creator                users.User
	symbol                 string
	crOn                   *time.Time
	lstUpOn                *time.Time
}

func createTokenBuilder(metaDataBuilderFactory metadata.BuilderFactory) tokens.TokenBuilder {
	out := tokenBuilder{
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:      nil,
		met:     nil,
		creator: nil,
		symbol:  "",
		crOn:    nil,
		lstUpOn: nil,
	}

	return &out
}

// Create initializes the token builder
func (build *tokenBuilder) Create() tokens.TokenBuilder {
	build.id = nil
	build.met = nil
	build.creator = nil
	build.symbol = ""
	build.crOn = nil
	build.lstUpOn = nil
	return build
}

// WithID adds an ID to the token builder
func (build *tokenBuilder) WithID(id *uuid.UUID) tokens.TokenBuilder {
	build.id = id
	return build
}

// WithMetaData adds metadata to the token builder
func (build *tokenBuilder) WithMetaData(met metadata.MetaData) tokens.TokenBuilder {
	build.met = met
	return build
}

// WithCreator adds a user creator to the token builder
func (build *tokenBuilder) WithCreator(creator users.User) tokens.TokenBuilder {
	build.creator = creator
	return build
}

// WithSymbol adds a symbol to the token builder
func (build *tokenBuilder) WithSymbol(symbol string) tokens.TokenBuilder {
	build.symbol = symbol
	return build
}

// CreatedOn adds a creation time to the token builder
func (build *tokenBuilder) CreatedOn(crOn time.Time) tokens.TokenBuilder {
	build.crOn = &crOn
	return build
}

// LastUpdatedOn adds a last updated on time to the token builder
func (build *tokenBuilder) LastUpdatedOn(lstUpOn time.Time) tokens.TokenBuilder {
	build.lstUpOn = &lstUpOn
	return build
}

// Now builds a new Token instance
func (build *tokenBuilder) Now() (tokens.Token, error) {

	if build.creator == nil {
		return nil, errors.New("the creator is mandatory in order to build a token instance")
	}

	if build.symbol == "" {
		return nil, errors.New("the symbol is mandatory in order to build a token instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a Token instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a Token instance")
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

	out := createToken(build.met.(*concrete_metadata.MetaData), build.creator.(*concrete_users.User), build.symbol)
	return out, nil
}
