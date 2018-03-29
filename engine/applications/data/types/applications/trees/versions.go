package trees

// Versions represents a list of versions
type Versions struct {
	Vers []*Version `json:"versions"`
}

// CreateVersions creates a Versions instance
func CreateVersions(vers []*Version) *Versions {
	out := Versions{
		Vers: vers,
	}

	return &out
}
