package commands

// Update represents an update command
type Update interface {
	GetOriginalJS() []byte
	GetNewJS() []byte
}
