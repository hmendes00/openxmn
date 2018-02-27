package objects

// Wallet represents a wallet
type Wallet struct {
	Balance map[string]*Token `json:"balance"`
	Stake   map[string]*Stake `json:"stake"`
}
