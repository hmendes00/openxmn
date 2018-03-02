package domain

// MetaDataBuilderFactory represents the metadata builder factory
type MetaDataBuilderFactory interface {
	Create() MetaDataBuilder
}
