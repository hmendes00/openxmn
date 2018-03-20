package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"

	cryptography "github.com/XMNBlockchain/exmachina-network/engine/domain/cryptography"
)

type signatureBuilder struct {
	publicKeyBuilderFactory cryptography.PublicKeyBuilderFactory
	pk                      *rsa.PrivateKey
	publicKey               cryptography.PublicKey
	pub                     *rsa.PublicKey
	encodedPub              string
	data                    []byte
	urlValues               url.Values
	v                       interface{}
	sig                     []byte
	encodedSig              string
}

func createSignatureBuilder(publicKeyBuilderFactory cryptography.PublicKeyBuilderFactory) cryptography.SignatureBuilder {
	out := signatureBuilder{
		publicKeyBuilderFactory: publicKeyBuilderFactory,
		pk:         nil,
		publicKey:  nil,
		pub:        nil,
		encodedPub: "",
		data:       nil,
		urlValues:  nil,
		v:          nil,
		sig:        nil,
		encodedSig: "",
	}
	return &out
}

// Create initializes the SignatureBuilder instance
func (build *signatureBuilder) Create() cryptography.SignatureBuilder {
	build.pk = nil
	build.publicKey = nil
	build.pub = nil
	build.encodedPub = ""
	build.data = nil
	build.urlValues = nil
	build.v = nil
	build.sig = nil
	build.encodedSig = ""
	return build
}

// WithPrivateKey adds a PrivateKey to the SignatureBuilder instance
func (build *signatureBuilder) WithPrivateKey(pk *rsa.PrivateKey) cryptography.SignatureBuilder {
	build.pk = pk
	return build
}

// WithPublicKey adds a PublicKey to the SignatureBuilder instance
func (build *signatureBuilder) WithPublicKey(pub *rsa.PublicKey) cryptography.SignatureBuilder {
	build.pub = pub
	return build
}

// WithEncodedPublicKey adds an encoded (base64) PublicKey to the SignatureBuilder instance
func (build *signatureBuilder) WithEncodedPublicKey(encodedPub string) cryptography.SignatureBuilder {
	build.encodedPub = encodedPub
	return build
}

// WithData adds data to the SignatureBuilder instance
func (build *signatureBuilder) WithData(data []byte) cryptography.SignatureBuilder {
	build.data = data
	return build
}

// WithURLValues adds urlValues to the SignatureBuilder instance
func (build *signatureBuilder) WithURLValues(urlValues url.Values) cryptography.SignatureBuilder {
	build.urlValues = urlValues
	return build
}

// WithInterface adds an instance to the SignatureBuilder instance
func (build *signatureBuilder) WithInterface(v interface{}) cryptography.SignatureBuilder {
	build.v = v
	return build
}

// WithSignature adds signature data to the SignatureBuilder instance
func (build *signatureBuilder) WithSignature(sig []byte) cryptography.SignatureBuilder {
	build.sig = sig
	return build
}

// WithSignature adds an encoded signature (base64) to the SignatureBuilder instance
func (build *signatureBuilder) WithEncodedSignature(encodedSig string) cryptography.SignatureBuilder {
	build.encodedSig = encodedSig
	return build
}

// Now builds a new Signature instance
func (build *signatureBuilder) Now() (cryptography.Signature, error) {

	if build.encodedSig != "" {
		jsSigAsByes, jsSigAsByesErr := base64.StdEncoding.DecodeString(build.encodedSig)
		if jsSigAsByesErr != nil {
			return nil, jsSigAsByesErr
		}

		out := new(Signature)
		jsErr := json.Unmarshal(jsSigAsByes, out)
		if jsErr != nil {
			return nil, jsErr
		}

		return out, nil
	}

	if build.urlValues != nil {
		str := build.urlValues.Encode()
		build.data = []byte(str)
	}

	if build.pk != nil {
		build.pub = &build.pk.PublicKey
	}

	if build.pub != nil && build.encodedPub != "" {
		return nil, errors.New("the PublicKey or the encoded PublicKey must be set, not both")
	}

	if build.pub == nil && build.encodedPub == "" {
		return nil, errors.New("the PublicKey or the encoded PublicKey is mandatory in order to build a Signature instance")
	}

	if build.data != nil && build.v != nil {
		return nil, errors.New("the data or the interface must be set, not both")
	}

	if build.data == nil && build.v == nil {
		return nil, errors.New("the data or the interface is mandatory in order to build a Signature instance")
	}

	if build.sig != nil && build.encodedSig != "" {
		return nil, errors.New("the signature or the encoded signature must be set, not both")
	}

	if build.encodedPub != "" {
		pub, pubErr := build.publicKeyBuilderFactory.Create().Create().WithEncodedString(build.encodedPub).Now()
		if pubErr != nil {
			return nil, pubErr
		}

		build.publicKey = pub
	}

	if build.pub != nil {
		pub, pubErr := build.publicKeyBuilderFactory.Create().Create().WithKey(build.pub).Now()
		if pubErr != nil {
			return nil, pubErr
		}

		build.publicKey = pub
	}

	if build.v != nil {
		js, jsErr := json.Marshal(build.v)
		if jsErr != nil {
			return nil, jsErr
		}

		build.data = js
	}

	if build.sig == nil {
		hash := sha256.New()
		hash.Write(build.data)
		build.data = hash.Sum(nil)
	}

	if build.sig == nil && build.pk != nil {

		if len(build.data) <= 0 {
			return nil, errors.New("the data is mandatory to build a Signature using a PrivateKey")
		}

		hash := sha256.New()
		hash.Write(build.data)
		sum := hash.Sum(nil)

		signature, errSignature := build.pk.Sign(rand.Reader, sum, crypto.SHA256)
		if errSignature != nil {
			return nil, errSignature
		}

		build.sig = signature
	}

	if build.sig == nil && build.encodedSig == "" {
		return nil, errors.New("the signature or the encoded signature is mandatory in order to build a Signature instance")
	}

	out, outErr := createSignature(build.data, build.sig, build.publicKey)
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
