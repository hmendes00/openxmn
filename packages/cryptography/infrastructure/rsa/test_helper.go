package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"testing"
)

// CreatePublicKeyForTests creates a PublicKey for tests
func CreatePublicKeyForTests(t *testing.T) *PublicKey {
	//variables:
	bitSize := 4096
	reader := rand.Reader
	key, keyErr := rsa.GenerateKey(reader, bitSize)

	if keyErr != nil {
		t.Errorf("there was a problem while generating an rsa.PrivateKey")
	}

	//execute:
	pk := createPublicKey(&key.PublicKey)
	return pk.(*PublicKey)
}

// CreateSignatureForTests creates a Signature for tests
func CreateSignatureForTests(t *testing.T) *Signature {
	pk, pkErr := CreatePrivateKeyBuilderFactory().Create().Create().Now()
	if pkErr != nil {
		t.Errorf("the returned error was expected to be nil, errror returned")
	}

	pubKey := createPublicKey(&pk.GetKey().PublicKey)
	data := []byte("this is some data we want to sign")
	hData := sha256.New()
	hData.Write(data)
	sig, sigErr := pk.GetKey().Sign(rand.Reader, hData.Sum(nil), crypto.SHA256)
	if sigErr != nil {
		t.Errorf("the returned error was expected to be nil, errror returned")
	}

	//execute:
	signature, signatureErr := createSignature(data, sig, pubKey)
	if signatureErr != nil {
		t.Errorf("the returned error was expected to be nil, errror returned")
	}

	return signature.(*Signature)
}
