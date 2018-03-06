package domain

// AggregatedSignedTransactionsRepository represents an aggregated signed transactions repository
type AggregatedSignedTransactionsRepository interface {
	Retrieve(dirPath string) (AggregatedSignedTransactions, error)
}
