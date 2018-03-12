package projects

// SettingsBuilderFactory represents a settings builder factory
type SettingsBuilderFactory interface {
	Create() SettingsBuilder
}
