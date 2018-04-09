package wallets

import (
	"errors"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

type walletBuilder struct {
	met    metadata.MetaData
	owner  users.User
	tok    tokens.Token
	amount float64
}

func createWalletBuilder() wallets.WalletBuilder {
	out := walletBuilder{
		met:    nil,
		owner:  nil,
		tok:    nil,
		amount: 0,
	}

	return &out
}

// Create initializes the wallet builder
func (build *walletBuilder) Create() wallets.WalletBuilder {
	build.met = nil
	build.owner = nil
	build.tok = nil
	build.amount = 0
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

// Now builds a new wallet instance
func (build *walletBuilder) Now() (wallets.Wallet, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a wallet instance")
	}

	if build.owner == nil {
		return nil, errors.New("the owner is mandatory in order to build a wallet instance")
	}

	if build.tok == nil {
		return nil, errors.New("the token is mandatory in order to build a wallet instance")
	}

	if build.amount <= 0 {
		return nil, errors.New("the amount must be greater than 0 in order to build a wallet instance")
	}

	out := createWallet(build.met.(*concrete_metadata.MetaData), build.owner.(*concrete_users.User), build.tok.(*concrete_tokens.Token), build.amount)
	return out, nil
}
