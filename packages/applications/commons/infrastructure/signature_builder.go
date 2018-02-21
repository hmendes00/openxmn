package infrastructure

import (
	"errors"
	"fmt"
	"net/http"

	commons "github.com/XMNBlockchain/core/packages/applications/commons/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	uuid "github.com/satori/go.uuid"
)

type signatureBuilder struct {
	sigBuilderFactory users.SignatureBuilderFactory
	r                 *http.Request
}

func createSignatureBuilder(sigBuilderFactory users.SignatureBuilderFactory) commons.SignatureBuilder {
	out := signatureBuilder{
		sigBuilderFactory: sigBuilderFactory,
		r:                 nil,
	}

	return &out
}

// Create initializes the SignatureBuilder instance
func (build *signatureBuilder) Create() commons.SignatureBuilder {
	build.r = nil
	return build
}

// WithRequest adds an http Request to the SignatureBuilder instance
func (build *signatureBuilder) WithRequest(r *http.Request) commons.SignatureBuilder {
	build.r = r
	return build
}

// Now builds a Signature instance
func (build *signatureBuilder) Now() (commons.Signature, error) {

	//parse form:
	pErr := build.r.ParseForm()
	if pErr != nil {
		str := fmt.Sprintf("there was an error while parsing form: %s", pErr.Error())
		return nil, errors.New(str)
	}

	//if there is a walletID:
	userIDAsString := build.r.Header.Get("user_id")
	if userIDAsString != "" {
		userID, userIDErr := uuid.FromString(userIDAsString)
		if userIDErr != nil {
			str := fmt.Sprintf("the user_id (%s) in the header is not a valid uuid: %s", userIDAsString, userIDErr.Error())
			return nil, errors.New(str)
		}

		//retrieve the encoded signature:
		encodedSignature := build.r.Header.Get("signature")

		userSig, userSigErr := build.sigBuilderFactory.Create().Create().WithEncodedSignature(encodedSignature).WithUserID(&userID).Now()
		if userSigErr != nil {
			str := fmt.Sprintf("there was an error while building the wallet signature: %s", userSigErr.Error())
			return nil, errors.New(str)
		}

		out := createSignatureWithUserSignature(userSig)
		return out, nil
	}

	out := createSignature()
	return out, nil
}
