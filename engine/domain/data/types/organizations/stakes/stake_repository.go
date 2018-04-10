package stakes

// StakeRepository represents a stake repository
type StakeRepository interface {
	Retrieve(dirPath string) (Stake, error)
}
