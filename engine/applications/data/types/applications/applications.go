package applications

// Applications represents a list of applications
type Applications struct {
	Apps *Application `json:"applications"`
}

// CreateApplications returns an applications instance
func CreateApplications(apps *Application) *Applications {
	out := Applications{
		Apps: apps,
	}

	return &out
}
