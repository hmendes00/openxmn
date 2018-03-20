package tables

// Table represents a table
type Table interface {
	GetSchema() Schema
	GetData() map[string]string
}
