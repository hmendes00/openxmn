package users

// UserRepository represents a stored user repository
type UserRepository interface {
	Retrieve(dirPath string) (User, error)
}
