package commands

// InsertBuilder represents an insert builder
type InsertBuilder interface {
	Create() InsertBuilder
	WithJS(js []byte) InsertBuilder
	Now() (Insert, error)
}
