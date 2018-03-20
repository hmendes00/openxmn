package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
)

// PublicKey represents a concrete PublicKey
type PublicKey struct {
	Key *rsa.PublicKey
}

type jsonifyPublicKey struct {
	EncodedKey string `json:"key"`
}

func createPublicKey(key *rsa.PublicKey) cryptography.PublicKey {
	out := PublicKey{
		Key: key,
	}

	return &out
}

func createPublicKeyFromEncodedString(encodedStr string) (cryptography.PublicKey, error) {
	block, _ := pem.Decode([]byte(encodedStr))
	if block == nil {
		return nil, errors.New("could not decode the public key")
	}

	pub, errPub := x509.ParsePKIXPublicKey(block.Bytes)
	if errPub != nil {
		return nil, errors.New("failed to parse DER encoded public key")
	}

	switch pubKey := pub.(type) {
	case *rsa.PublicKey:
		out := createPublicKey(pubKey)
		return out, nil
	default:
		return nil, errors.New("the given public key is not of type: *rsa.PublicKey")
	}
}

// String returns the string representaton of the public key
func (pub *PublicKey) String() (string, error) {
	public, publicErr := x509.MarshalPKIXPublicKey(pub.Key)
	if publicErr != nil {
		return "", publicErr
	}

	encodedStr := base64.StdEncoding.EncodeToString(public)
	strAsRunes := []rune(encodedStr)
	splitEncodedStr := ""
	for index, oneChar := range strAsRunes {
		if index%64 == 0 && index != 0 {
			splitEncodedStr = fmt.Sprintf("%s\n%c", splitEncodedStr, oneChar)
			continue
		}

		splitEncodedStr = fmt.Sprintf("%s%c", splitEncodedStr, oneChar)
	}

	return fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----\n", splitEncodedStr), nil
}

// GetKey returns the *rsa.PublicKey associated with the public key
func (pub *PublicKey) GetKey() *rsa.PublicKey {
	return pub.Key
}

// Compare returns true if the given PublicKey is the same as the current PublicKey
func (pub *PublicKey) Compare(pk cryptography.PublicKey) bool {
	self, selfErr := pub.String()
	if selfErr != nil {
		return false
	}

	pkStr, pkStrErr := pk.String()
	if pkStrErr != nil {
		return false
	}

	return self == pkStr
}

// MarshalJSON transform a PublicKey to JSON
func (pub *PublicKey) MarshalJSON() ([]byte, error) {
	str, strErr := pub.String()
	if strErr != nil {
		return nil, strErr
	}

	jsonify := jsonifyPublicKey{
		EncodedKey: str,
	}

	js, jsErr := json.Marshal(jsonify)
	if jsErr != nil {
		return nil, jsErr
	}

	return js, nil
}

// UnmarshalJSON transform the data to a PublicKey instance
func (pub *PublicKey) UnmarshalJSON(data []byte) error {

	jsonify := new(jsonifyPublicKey)
	unErr := json.Unmarshal(data, &jsonify)
	if unErr != nil {
		return unErr
	}

	pubKey, pubKeyErr := createPublicKeyFromEncodedString(jsonify.EncodedKey)
	if pubKeyErr != nil {
		return pubKeyErr
	}

	pub.Key = pubKey.GetKey()
	return nil
}
