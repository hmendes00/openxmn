package domain

// MetaDataBuilderFactory represents a chained block metadata builder factory
type MetaDataBuilderFactory interface {
	Create() MetaDataBuilder
}
