package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"reflect"
	"testing"
)

func TestCreateCipher_thenDecipher_Success(t *testing.T) {

	//variables:
	text := []byte("here is some data")
	lbl := []byte("some lbl")
	hash := sha256.New()
	pk, _ := CreatePrivateKeyBuilderFactoryForTests().Create().Create().Now()
	pubKey := pk.GetPublicKey()
	cipherText, _ := rsa.EncryptOAEP(hash, rand.Reader, pubKey.GetKey(), text, lbl)
	sig, _ := CreateSignatureBuilderFactoryForTests().Create().Create().WithData(text).WithPrivateKey(pk.GetKey()).Now()

	//execute:
	ci := createCipher(hash, cipherText, lbl, sig.(*Signature))

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
