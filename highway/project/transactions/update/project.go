package update

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Project represents an update project transaction
type Project struct {
	ProjectID                    *uuid.UUID    `json:"project_id"`
	PricePerTrx                  float64       `json:"price_per_transaction"`
	ShareToProcessors            float64       `json:"share_to_processors"`
	ShareToLeaders               float64       `json:"share_to_leaders"`
	ShareToVerifiers             float64       `json:"share_to_verifiers"`
	ShareToBlocker               float64       `json:"share_to_blocker"`
	ShareToShareHolders          float64       `json:"share_to_shareholders"`
	AmountOfQuotasNeededPerBlock uint          `json:"amount_of_quotas_needed_per_block"`
	BlockDuration                time.Duration `json:"block_duration"`
}
