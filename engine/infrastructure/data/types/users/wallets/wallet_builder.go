package wallets

import (
	"errors"
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

type walletBuilder struct {
	metaDataBuilderFactory metadata.BuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	owner                  users.User
	tok                    tokens.Token
	amount                 float64
	crOn                   *time.Time
	lstUpOn                *time.Time
}

func createWalletBuilder(metaDataBuilderFactory metadata.BuilderFactory) wallets.WalletBuilder {
	out := walletBuilder{
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:      nil,
		met:     nil,
		owner:   nil,
		tok:     nil,
		amount:  0,
		crOn:    nil,
		lstUpOn: nil,
	}

	return &out
}

// Create initializes the wallet builder
func (build *walletBuilder) Create() wallets.WalletBuilder {
	build.id = nil
	build.met = nil
	build.owner = nil
	build.tok = nil
	build.amount = 0
	build.crOn = nil
	build.lstUpOn = nil
	return build
}

// WithID adds an ID to the wallet builder
func (build *walletBuilder) WithID(id *uuid.UUID) wallets.WalletBuilder {
	build.id = id
	return build
}

// WithMetaData adds metadata to the wallet builder
func (build *walletBuilder) WithMetaData(met metadata.MetaData) wallets.WalletBuilder {
	build.met = met
	return build
}

// WithOwner adds an owner to the wallet builder
func (build *walletBuilder) WithOwner(owner users.User) wallets.WalletBuilder {
	build.owner = owner
	return build
}

// WithToken adds a token to the wallet builder
func (build *walletBuilder) WithToken(tok tokens.Token) wallets.WalletBuilder {
	build.tok = tok
	return build
}

// WithAmount adds an amount to the wallet builder
func (build *walletBuilder) WithAmount(amount float64) wallets.WalletBuilder {
	build.amount = amount
	return build
}

// CreatedOn adds a creation time to the wallet builder
func (build *walletBuilder) CreatedOn(crOn time.Time) wallets.WalletBuilder {
	build.crOn = &crOn
	return build
}

// LastUpdatedOn adds an updated on time to the wallet builder
func (build *walletBuilder) LastUpdatedOn(lstUpOn time.Time) wallets.WalletBuilder {
	build.lstUpOn = &lstUpOn
	return build
}

// Now builds a new wallet instance
func (build *walletBuilder) Now() (wallets.Wallet, error) {

	if build.owner == nil {
		return nil, errors.New("the owner is mandatory in order to build a wallet instance")
	}

	if build.tok == nil {
		return nil, errors.New("the token is mandatory in order to build a wallet instance")
	}

	if build.amount <= 0 {
		return nil, errors.New("the amount must be greater than 0 in order to build a wallet instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a Wallet instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a Wallet instance")
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

	out := createWallet(build.met.(*concrete_metadata.MetaData), build.owner.(*concrete_users.User), build.tok.(*concrete_tokens.Token), build.amount)
	return out, nil
}
