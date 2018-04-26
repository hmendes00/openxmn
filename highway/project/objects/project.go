package objects

import (
	"errors"
	"fmt"
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
)

// Project represents a project
type Project struct {
	Met                          *concrete_metadata.MetaData `json:"metadata"`
	RelPath                      string                      `json:"relative_path"`
	Org                          *Organization               `json:"organization"`
	PricePerTrx                  float64                     `json:"price_per_transaction"`
	ShareToProcessors            float64                     `json:"share_to_processors"`
	ShareToLeaders               float64                     `json:"share_to_leaders"`
	ShareToVerifiers             float64                     `json:"share_to_verifiers"`
	ShareToBlocker               float64                     `json:"share_to_blocker"`
	ShareToShareHolders          float64                     `json:"share_to_shareholders"`
	AmountOfQuotasNeededPerBlock uint                        `json:"amount_of_quotas_needed_per_block"`
	BlockDuration                time.Duration               `json:"block_duration"`
}

// CreateProject creates a new Project instance
func CreateProject(
	met metadata.MetaData,
	relPath string,
	org *Organization,
	pricePerTrx float64,
	shareToProcessors float64,
	shareToLeaders float64,
	shareToVerifiers float64,
	shareToBlocker float64,
	shareToShareHolders float64,
	amountOfQuotasNeededPerBlock uint,
	blockDuration time.Duration,
) (*Project, error) {

	allShares := shareToProcessors + shareToLeaders + shareToVerifiers + shareToBlocker + shareToShareHolders
	if allShares != float64(1) {
		str := fmt.Sprintf("the cumulative shares does not equal 1.  It equals: %f", allShares)
		return nil, errors.New(str)
	}

	out := Project{
		Met:                          met.(*concrete_metadata.MetaData),
		RelPath:                      relPath,
		Org:                          org,
		PricePerTrx:                  pricePerTrx,
		ShareToProcessors:            shareToProcessors,
		ShareToLeaders:               shareToLeaders,
		ShareToVerifiers:             shareToVerifiers,
		ShareToBlocker:               shareToBlocker,
		ShareToShareHolders:          shareToShareHolders,
		AmountOfQuotasNeededPerBlock: amountOfQuotasNeededPerBlock,
		BlockDuration:                blockDuration,
	}

	return &out, nil
}
