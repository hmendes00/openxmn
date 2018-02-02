package custom

import (
	"errors"

	custom "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/custom"
)

type customBuilder struct {
	cr custom.Create
}

func createCustomBuilder() custom.Builder {
	out := customBuilder{
		cr: nil,
	}

	return &out
}

// Create creates a new Builder instance
func (build *customBuilder) Create() custom.Builder {
	build.cr = nil
	return build
}

// WithCreate adds a Create instance to the custom builder
func (build *customBuilder) WithCreate(cr custom.Create) custom.Builder {
	build.cr = cr
	return build
}

// Now builds a new Custom instance
func (build *customBuilder) Now() (custom.Custom, error) {
	if build.cr == nil {
		return nil, errors.New("there must 1 custom transaction, none given")
	}

	out := createCustomWithCreate(build.cr.(*Create))
	return out, nil
}
