package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
)

// CreatePublicKeyForTests creates a PublicKey for tests
func CreatePublicKeyForTests() *PublicKey {
	//variables:
	bitSize := 4096
	reader := rand.Reader
	key, _ := rsa.GenerateKey(reader, bitSize)

	//execute:
	pk := createPublicKey(&key.PublicKey)
	return pk.(*PublicKey)
}

// CreateSignatureForTests creates a Signature for tests
func CreateSignatureForTests() *Signature {
	pk, _ := CreatePrivateKeyBuilderFactory().Create().Create().Now()

	pubKey := createPublicKey(&pk.GetKey().PublicKey)
	data := []byte("this is some data we want to sign")
	hData := sha256.New()
	hData.Write(data)
	sig, _ := pk.GetKey().Sign(rand.Reader, hData.Sum(nil), crypto.SHA256)

	//execute:
	signature, _ := createSignature(data, sig, pubKey)
	return signature.(*Signature)
}

// CreatePublicKeyBuilderFactoryForTests creates a new PublicKeyBuilderFactory for tests
func CreatePublicKeyBuilderFactoryForTests() cryptography.PublicKeyBuilderFactory {
	out := CreatePublicKeyBuilderFactory()
	return out
}

// CreateSignatureBuilderFactoryForTests creates a new SignatureBuilderFactory for tests
func CreateSignatureBuilderFactoryForTests() cryptography.SignatureBuilderFactory {
	publicKeyBuilderFactory := CreatePublicKeyBuilderFactoryForTests()
	out := CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	return out
}

// CreatePrivateKeyBuilderFactoryForTests creates a new PrivateKeyBuilderFactory for tests
func CreatePrivateKeyBuilderFactoryForTests() cryptography.PrivateKeyBuilderFactory {
	out := CreatePrivateKeyBuilderFactory()
	return out
}
