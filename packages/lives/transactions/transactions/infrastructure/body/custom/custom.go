package custom

import (
	custom "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/custom"
)

// Custom represents the concrete custom transaction
type Custom struct {
	Cr *Create `json:"create"`
}

func createCustomWithCreate(cr *Create) custom.Custom {
	out := Custom{
		Cr: cr,
	}

	return &out
}

// HasCreate returns true if there is a Create instance, false otherwise
func (cu *Custom) HasCreate() bool {
	return cu.Cr != nil
}

// GetCreate returns the Create instance
func (cu *Custom) GetCreate() custom.Create {
	return cu.Cr
}
