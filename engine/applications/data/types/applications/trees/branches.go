package trees

// Branches represents a list of versions
type Branches struct {
	Brchs []*Branch `json:"branches"`
}

// CreateBranches creates a Branches instance
func CreateBranches(brchs []*Branch) *Branches {
	out := Branches{
		Brchs: brchs,
	}

	return &out
}
