package domain

// MetaDataBuilderFactory represents a MetaDataBuilder factory
type MetaDataBuilderFactory interface {
	Create() MetaDataBuilder
}
