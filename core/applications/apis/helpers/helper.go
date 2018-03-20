package helpers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/users"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/users"
)

// FromRequestToUserSignature converts an HTTP Request to a User Signature
func FromRequestToUserSignature(r *http.Request) (users.Signature, error) {
	//parse form:
	pErr := r.ParseForm()
	if pErr != nil {
		return nil, pErr
	}

	//if there is a signature:
	sigAsString := r.Header.Get("signature")
	if sigAsString != "" {
		js, decErr := base64.StdEncoding.DecodeString(sigAsString)
		if decErr != nil {
			return nil, decErr
		}

		sig := new(concrete_users.Signature)
		jsErr := json.Unmarshal(js, sig)
		if jsErr != nil {
			return nil, jsErr
		}

		return sig, nil
	}

	return nil, nil
}
