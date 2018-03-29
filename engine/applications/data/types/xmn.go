package types

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/remote"
)

// XMN represents the core data type
type XMN struct {
	apps          *remote.Files
	servers       *remote.Files
	organizations *remote.Files
	safes         *remote.Files
	stakes        *remote.Files
	tokens        *remote.Files
	wallets       *remote.Files
}

// CreateXMN creates a new XMN instance
func CreateXMN(
	apps *remote.Files,
	servers *remote.Files,
	organizations *remote.Files,
	safes *remote.Files,
	stakes *remote.Files,
	tokens *remote.Files,
	wallets *remote.Files,
) *XMN {
	out := XMN{
		apps:          apps,
		servers:       servers,
		organizations: organizations,
		safes:         safes,
		stakes:        stakes,
		tokens:        tokens,
		wallets:       wallets,
	}

	return &out
}
