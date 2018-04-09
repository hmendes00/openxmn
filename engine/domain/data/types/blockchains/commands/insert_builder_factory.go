package commands

// InsertBuilderFactory represents an insert builder factory
type InsertBuilderFactory interface {
	Create() InsertBuilder
}
