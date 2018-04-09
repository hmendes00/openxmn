package types

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/remote"
	blkmetadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
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

// PayValidators pays validators of a block
func (core *XMN) PayValidators(validatedBlkMetaData blkmetadata.MetaData, sigs users.Signatures) error {
	return nil
}

// PayBlockBuilder pays the builder of a block
func (core *XMN) PayBlockBuilder(signedBlkMetaData blkmetadata.MetaData, sig users.Signature) error {
	return nil
}

// PayLeader pays the leader that aggregated transactions
func (core *XMN) PayLeader(aggregatedTrsMetaData blkmetadata.MetaData, sig users.Signature) error {
	return nil
}

// PayAtomicTransactionProcessor pays the atomic transaction processor
func (core *XMN) PayAtomicTransactionProcessor(atomicTrsMetaData blkmetadata.MetaData, sig users.Signature) error {
	return nil
}

// PayTransactionProcessor pays the transaction processor
func (core *XMN) PayTransactionProcessor(trsMetaData blkmetadata.MetaData, sig users.Signature) error {
	return nil
}
