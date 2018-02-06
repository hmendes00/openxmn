package servers

// DeleteBuilderFactory represents the builder factory of a delete server transaction
type DeleteBuilderFactory interface {
	Create() DeleteBuilder
}
