package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateSignature_Success(t *testing.T) {

	//variables:
	pk, pkErr := CreatePrivateKeyBuilderFactory().Create().Create().Now()
	if pkErr != nil {
		t.Errorf("the returned error was expected to be nil, errror returned: %s", pkErr.Error())
	}

	pubKey := createPublicKey(&pk.GetKey().PublicKey)
	data := []byte("this is some data we want to sign")
	hData := sha256.New()
	hData.Write(data)
	sig, sigErr := pk.GetKey().Sign(rand.Reader, hData.Sum(nil), crypto.SHA256)
	if sigErr != nil {
		t.Errorf("the returned error was expected to be nil, errror returned: %s", sigErr.Error())
	}

	//execute:
	signature, signatureErr := createSignature(data, sig, pubKey)
	if signatureErr != nil {
		t.Errorf("the returned error was expected to be nil, errror returned: %s", signatureErr.Error())
	}

	if signature == nil {
		t.Errorf("the returned instance was expected to be an instance, nil returned")
	}
}

func TestCreateSignature_convertToJSON_backAndForth_Success(t *testing.T) {

	//variables:
	empty := new(Signature)
	sig := CreateSignatureForTests()

	//execute:
	convert.ConvertToJSON(t, sig, empty)
}
