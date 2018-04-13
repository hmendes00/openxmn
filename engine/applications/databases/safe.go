package databases

import (
	"errors"
	"fmt"

	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	uuid "github.com/satori/go.uuid"
)

// Safe represents a safe database
type Safe struct {
	sfs map[string]safes.Safe
}

// CreateSafe creates a new Safe instance
func CreateSafe() *Safe {
	out := Safe{
		sfs: map[string]safes.Safe{},
	}

	return &out
}

// RetrieveByID retrieves the safe by ID
func (db *Safe) RetrieveByID(id *uuid.UUID) (safes.Safe, error) {
	idAsString := id.String()
	if oneSafe, ok := db.sfs[idAsString]; ok {
		return oneSafe, nil
	}

	str := fmt.Sprintf("the safe (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

func (db *Safe) insert(saf safes.Safe) error {
	id := saf.GetMetaData().GetID()
	idAsString := id.String()
	_, retSafeErr := db.RetrieveByID(id)
	if retSafeErr == nil {
		str := fmt.Sprintf("there is already a safe with ID: %s", idAsString)
		return errors.New(str)
	}

	db.sfs[idAsString] = saf
	return nil
}

// Update updates a safe
func (db *Safe) Update(old safes.Safe, new safes.Safe) error {
	newSafeID := new.GetMetaData().GetID()
	newSafeIDAsString := newSafeID.String()
	_, retNewSafeErr := db.RetrieveByID(newSafeID)
	if retNewSafeErr == nil {
		str := fmt.Sprintf("the new safe (ID: %s) already exists", newSafeIDAsString)
		return errors.New(str)
	}

	delErr := db.Delete(old)
	if delErr != nil {
		return delErr
	}

	insErr := db.insert(new)
	if insErr != nil {
		return insErr
	}

	return nil
}

// Delete deletes a safe
func (db *Safe) Delete(safe safes.Safe) error {
	id := safe.GetMetaData().GetID()
	_, retSafeErr := db.RetrieveByID(id)
	if retSafeErr != nil {
		return retSafeErr
	}

	idAsString := id.String()
	delete(db.sfs, idAsString)
	return nil
}
