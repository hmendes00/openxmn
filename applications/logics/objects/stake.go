package objects

// Stake represents tokens at stake to an organization
type Stake struct {
	Balance map[string]*Token `json:"balance"`
	Org     *Organization     `json:"organization"`
}
