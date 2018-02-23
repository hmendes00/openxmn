package users

import (
	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
)

// User represents the concrete server transaction
type User struct {
	CR  *Save   `json:"save"`
	DEL *Delete `json:"delete"`
}

func createUserWithSave(cr *Save) users.User {
	out := User{
		CR:  cr,
		DEL: nil,
	}

	return &out
}

func createUserWithDelete(del *Delete) users.User {
	out := User{
		CR:  nil,
		DEL: del,
	}

	return &out
}

// HasSave returns true if there is a Save, false otherwise
func (serv *User) HasSave() bool {
	return serv.CR != nil
}

// GetSave returns the Save transaction
func (serv *User) GetSave() users.Save {
	return serv.CR
}

// HasDelete returns true if there is a Delete, false otherwise
func (serv *User) HasDelete() bool {
	return serv.DEL != nil
}

// GetDelete returns the Delete transaction
func (serv *User) GetDelete() users.Delete {
	return serv.DEL
}
