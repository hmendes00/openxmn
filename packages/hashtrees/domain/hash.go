package domain

// Hash represents a single hash
type Hash interface {
	String() string
	Get() []byte
	Compare(h Hash) bool
}
