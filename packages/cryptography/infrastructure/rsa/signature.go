package rsa

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"hash"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

// Signature represents a cryptographic signature
type Signature struct {
	hash hash.Hash
	data []byte
	sig  []byte
	pub  cryptography.PublicKey
}

type jsonifySignature struct {
	Data   string `json:"data"`
	Sig    string `json:"signature"`
	PubKey string `json:"public_key"`
}

func createSignature(data []byte, sig []byte, pub cryptography.PublicKey) (cryptography.Signature, error) {
	//creates the hash:
	hash := sha256.New()
	hash.Write(data)

	out := Signature{
		hash: hash,
		data: data,
		sig:  sig,
		pub:  pub,
	}

	//make sure the signture works:
	if !out.verify() {
		return nil, errors.New("failed to validate the authenticity of the signature")
	}

	return &out, nil
}

func (sig *Signature) verify() bool {
	sum := sig.hash.Sum(nil)
	pubKey := sig.pub.GetKey()
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, sum, sig.sig)
	if err != nil {
		return false
	}
	return true
}

// String encodes a base64 representation of the signature
func (sig *Signature) String() string {
	js, jsErr := json.Marshal(sig)
	if jsErr != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(js)
}

// GetPublicKey returns the PublicKey
func (sig *Signature) GetPublicKey() cryptography.PublicKey {
	return sig.pub
}

// MarshalJSON transform a Signature to JSON
func (sig *Signature) MarshalJSON() ([]byte, error) {
	encodedData := base64.StdEncoding.EncodeToString(sig.data)
	encodedSig := base64.StdEncoding.EncodeToString(sig.sig)
	encodedPublicKey, encodedPublicKeyErr := sig.pub.String()
	if encodedPublicKeyErr != nil {
		return nil, encodedPublicKeyErr
	}

	jsonify := jsonifySignature{
		Data:   encodedData,
		Sig:    encodedSig,
		PubKey: encodedPublicKey,
	}

	js, jsErr := json.Marshal(jsonify)
	if jsErr != nil {
		return nil, jsErr
	}

	return js, nil
}

// UnmarshalJSON transform the data to a Signature instance
func (sig *Signature) UnmarshalJSON(data []byte) error {

	jsonify := new(jsonifySignature)
	unErr := json.Unmarshal(data, &jsonify)
	if unErr != nil {
		return unErr
	}

	data, dataErr := base64.StdEncoding.DecodeString(jsonify.Data)
	if dataErr != nil {
		return dataErr
	}

	decSig, decSigErr := base64.StdEncoding.DecodeString(jsonify.Sig)
	if decSigErr != nil {
		return decSigErr
	}

	pubKey, pubKeyErr := createPublicKeyFromEncodedString(jsonify.PubKey)
	if pubKeyErr != nil {
		return pubKeyErr
	}

	hash := sha256.New()
	hash.Write(data)

	sig.data = data
	sig.hash = hash
	sig.sig = decSig
	sig.pub = pubKey

	if !sig.verify() {
		return errors.New("failed to validate the authenticity of the signature")
	}

	return nil
}
