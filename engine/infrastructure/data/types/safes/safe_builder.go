package safes

import (
	"errors"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
)

type safeBuilder struct {
	met  metadata.MetaData
	tok  tokens.Token
	ciph cryptography.Cipher
}

func createSafeBuilder() safes.SafeBuilder {
	out := safeBuilder{
		met:  nil,
		tok:  nil,
		ciph: nil,
	}

	return &out
}

// Create initializes the safe
func (build *safeBuilder) Create() safes.SafeBuilder {
	build.met = nil
	build.tok = nil
	build.ciph = nil
	return build
}

// WithMetaData adds metadata to the SafeBuilder instance
func (build *safeBuilder) WithMetaData(met metadata.MetaData) safes.SafeBuilder {
	build.met = met
	return build
}

// WithToken adds a token to the SafeBuilder instance
func (build *safeBuilder) WithToken(tok tokens.Token) safes.SafeBuilder {
	build.tok = tok
	return build
}

// WithCipher adds a cipher to the SafeBuilder instance
func (build *safeBuilder) WithCipher(cipher cryptography.Cipher) safes.SafeBuilder {
	build.ciph = cipher
	return build
}

// Now builds a new Safe instance
func (build *safeBuilder) Now() (safes.Safe, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a Safe instance")
	}

	if build.tok == nil {
		return nil, errors.New("the token is mandatory in order to build a Safe instance")
	}

	if build.ciph == nil {
		return nil, errors.New("the cipher is mandatory in order to build a Safe instance")
	}

	out := createSafe(build.met.(*concrete_metadata.MetaData), build.tok.(*concrete_tokens.Token), build.ciph.(*concrete_cryptography.Cipher))
	return out, nil
}
