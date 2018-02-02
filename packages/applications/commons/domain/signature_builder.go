package domain

import "net/http"

// SignatureBuilder represents a Signature builder
type SignatureBuilder interface {
	Create() SignatureBuilder
	WithRequest(r *http.Request) SignatureBuilder
	Now() (Signature, error)
}
