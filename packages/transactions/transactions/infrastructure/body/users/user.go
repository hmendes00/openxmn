package users

import (
	users "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/users"
)

// User represents the concrete server transaction
type User struct {
	CR  *Create `json:"create"`
	DEL *Delete `json:"delete"`
	UP  *Update `json:"update"`
}

func createUserWithCreate(cr *Create) users.User {
	out := User{
		CR:  cr,
		DEL: nil,
		UP:  nil,
	}

	return &out
}

func createUserWithDelete(del *Delete) users.User {
	out := User{
		CR:  nil,
		DEL: del,
		UP:  nil,
	}

	return &out
}

func createUserWithUpdate(up *Update) users.User {
	out := User{
		CR:  nil,
		DEL: nil,
		UP:  up,
	}

	return &out
}

// HasCreate returns true if there is a Create, false otherwise
func (serv *User) HasCreate() bool {
	return serv.CR != nil
}

// GetCreate returns the Create transaction
func (serv *User) GetCreate() users.Create {
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

// HasUpdate returns true if there is an Update, false otherwise
func (serv *User) HasUpdate() bool {
	return serv.UP != nil
}

// GetUpdate returns the Update transaction
func (serv *User) GetUpdate() users.Update {
	return serv.UP
}
