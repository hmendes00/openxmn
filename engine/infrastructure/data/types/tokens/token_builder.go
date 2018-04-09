package tokens

import (
	"errors"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

type tokenBuilder struct {
	met     metadata.MetaData
	creator users.User
	symbol  string
	amount  int
}

func createTokenBuilder() tokens.TokenBuilder {
	out := tokenBuilder{
		met:     nil,
		creator: nil,
		symbol:  "",
		amount:  0,
	}

	return &out
}

// Create initializes the token builder
func (build *tokenBuilder) Create() tokens.TokenBuilder {
	build.met = nil
	build.creator = nil
	build.symbol = ""
	build.amount = 0
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

// WithAmount adds an amount to the token builder
func (build *tokenBuilder) WithAmount(amount int) tokens.TokenBuilder {
	build.amount = amount
	return build
}

// Now builds a new Token instance
func (build *tokenBuilder) Now() (tokens.Token, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a token instance")
	}

	if build.creator == nil {
		return nil, errors.New("the creator is mandatory in order to build a token instance")
	}

	if build.symbol == "" {
		return nil, errors.New("the symbol is mandatory in order to build a token instance")
	}

	if build.amount <= 0 {
		return nil, errors.New("the amount must be greater than 0 in order to build a token instance")
	}

	out := createToken(build.met.(*concrete_metadata.MetaData), build.creator.(*concrete_users.User), build.symbol, build.amount)
	return out, nil
}
