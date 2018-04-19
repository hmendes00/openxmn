package commands

// Repository represents a commands repository
type Repository interface {
	Retrieve(dirPath string) (Commands, error)
}
