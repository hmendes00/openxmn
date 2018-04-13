package databases

import (
	"errors"
	"fmt"

	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// User represents a user database
type User struct {
	usrs map[string]users.User
}

// CreateUser creates a user database
func CreateUser() *User {
	out := User{
		usrs: map[string]users.User{},
	}

	return &out
}

// RetrieveByID retrieves a user by ID
func (db *User) RetrieveByID(id *uuid.UUID) (users.User, error) {
	idAsString := id.String()
	if oneUsr, ok := db.usrs[idAsString]; ok {
		return oneUsr, nil
	}

	str := fmt.Sprintf("the user (ID: %s) could not be found", idAsString)
	return nil, errors.New(str)
}

// Insert inserts a new user to the database
func (db *User) Insert(usr users.User) error {
	id := usr.GetMetaData().GetID()
	idAsString := id.String()
	_, retUsrErr := db.RetrieveByID(id)
	if retUsrErr == nil {
		str := fmt.Sprintf("there is already a user with ID: %s", idAsString)
		return errors.New(str)
	}

	db.usrs[idAsString] = usr
	return nil
}

// Update updates a user to the database
func (db *User) Update(old users.User, new users.User) error {
	newUsrID := new.GetMetaData().GetID()
	newUsrIDAsString := newUsrID.String()
	_, retNewUsrErr := db.RetrieveByID(newUsrID)
	if retNewUsrErr == nil {
		str := fmt.Sprintf("the new user (ID: %s) already exists", newUsrIDAsString)
		return errors.New(str)
	}

	delErr := db.Delete(old)
	if delErr != nil {
		return delErr
	}

	insErr := db.Insert(new)
	if insErr != nil {
		return insErr
	}

	return nil
}

// Delete deletes a user from the database
func (db *User) Delete(usr users.User) error {
	id := usr.GetMetaData().GetID()
	_, retUsrErr := db.RetrieveByID(id)
	if retUsrErr != nil {
		return retUsrErr
	}

	idAsString := id.String()
	delete(db.usrs, idAsString)
	return nil
}
