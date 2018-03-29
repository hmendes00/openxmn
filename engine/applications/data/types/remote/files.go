package remote

// Files represents a list of files
type Files struct {
	Fils []*File `json:"files"`
}

// CreateFiles creates a files instance
func CreateFiles(fils []*File) *Files {
	out := Files{
		Fils: fils,
	}

	return &out
}
