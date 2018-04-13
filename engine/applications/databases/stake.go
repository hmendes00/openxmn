package databases

import (
	"errors"
	"fmt"

	stakes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/stakes"
	uuid "github.com/satori/go.uuid"
)

// Stake represents a stake database
type Stake struct {
	stks map[string]stakes.Stake
}

// CreateStake creates a new Stake instance
func CreateStake() *Stake {
	out := Stake{
		stks: map[string]stakes.Stake{},
	}

	return &out
}

// RetrieveByID retrieves a Stake by ID
func (db *Stake) RetrieveByID(id *uuid.UUID) (stakes.Stake, error) {
	idAsString := id.String()
	if oneStk, ok := db.stks[idAsString]; ok {
		return oneStk, nil
	}

	str := fmt.Sprintf("the stake (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

// Insert insert a new Stake
func (db *Stake) Insert(stk stakes.Stake) error {
	id := stk.GetMetaData().GetID()
	idAsString := id.String()
	_, retStkErr := db.RetrieveByID(id)
	if retStkErr == nil {
		str := fmt.Sprintf("there is already a stake with ID: %s", idAsString)
		return errors.New(str)
	}

	db.stks[idAsString] = stk
	return nil
}

// Delete deletes a stake
func (db *Stake) Delete(stk stakes.Stake) error {
	id := stk.GetMetaData().GetID()
	_, retStkErr := db.RetrieveByID(id)
	if retStkErr != nil {
		return retStkErr
	}

	idAsString := id.String()
	delete(db.stks, idAsString)
	return nil
}
