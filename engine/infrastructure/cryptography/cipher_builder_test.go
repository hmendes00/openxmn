package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"reflect"
	"testing"
)

func TestBuildCipher_Success(t *testing.T) {

	//variables:
	text := []byte("here is some data")
	lbl := []byte("some lbl")
	pk, _ := CreatePrivateKeyBuilderFactoryForTests().Create().Create().Now()
	pubKey := pk.GetPublicKey()

	//execute:
	build := CreateCipherBuilderFactoryWithCustomPKForTests(pk)
	ci, ciErr := build.Create().Create().WithLabel(lbl).WithText(text).WithPublicKey(pubKey).Now()
	if ciErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", ciErr.Error())
	}

	retLbl := ci.GetLabel()
	retDecipheredTxt, _ := ci.Decipher(pk)

	if !reflect.DeepEqual(retLbl, lbl) {
		t.Errorf("the returned label is invalid.")
	}

	if !reflect.DeepEqual(retDecipheredTxt, text) {
		t.Errorf("the returned deciphered text is invalid.")
	}

}

func TestBuildCipher_WithCipherData_Success(t *testing.T) {

	//variables:
	text := []byte("here is some data")
	lbl := []byte("some lbl")
	hash := sha256.New()
	pk, _ := CreatePrivateKeyBuilderFactoryForTests().Create().Create().Now()
	pubKey := pk.GetPublicKey()
	cipherText, _ := rsa.EncryptOAEP(hash, rand.Reader, pubKey.GetKey(), text, lbl)
	sig, _ := CreateSignatureBuilderFactoryForTests().Create().Create().WithData(text).WithPrivateKey(pk.GetKey()).Now()

	//execute:
	build := CreateCipherBuilderFactoryForTests()
	ci, ciErr := build.Create().Create().WithCipherText(cipherText).WithHash(hash).WithLabel(lbl).WithPublicKey(pubKey).WithSignature(sig).Now()
	if ciErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", ciErr.Error())
	}

	retText := ci.GetText()
	retLbl := ci.GetLabel()
	retSig := ci.GetSignature()
	retDecipheredTxt, _ := ci.Decipher(pk)

	if !reflect.DeepEqual(retText, cipherText) {
		t.Errorf("the returned cipherText is invalid.")
	}

	if !reflect.DeepEqual(retLbl, lbl) {
		t.Errorf("the returned label is invalid.")
	}

	if !reflect.DeepEqual(retSig, sig) {
		t.Errorf("the returned signature is invalid.")
	}

	if !reflect.DeepEqual(retDecipheredTxt, text) {
		t.Errorf("the returned deciphered text is invalid.")
	}

}
