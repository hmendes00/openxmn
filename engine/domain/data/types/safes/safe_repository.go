package safes

// SafeRepository represents a safe repository
type SafeRepository interface {
	Retrieve(dirPath string) (Safe, error)
}
