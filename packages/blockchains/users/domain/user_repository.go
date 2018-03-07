package domain

// UserRepository represents a user repository
type UserRepository interface {
	Retrieve(dirPath string) (User, error)
}
