package users

import (
	"reflect"
	"testing"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	uuid "github.com/satori/go.uuid"
)

func TestBuildCreate_withID_withPublicKey_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests(t)

	//execute:
	build := createCreateBuilder()
	cr, crErr := build.Create().WithID(&id).WithPublicKey(pk).Now()

	if crErr != nil {
		t.Errorf("the returned error was expected to be nil, instance returned: %s", crErr.Error())
	}

	retID := cr.GetID()
	retPK := cr.GetPublicKey()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	if !reflect.DeepEqual(pk, retPK) {
		t.Errorf("the returned public key is invalid")
	}

}

func TestBuildCreate_withoutID_withPublicKey_returnsError(t *testing.T) {

	//execute:
	pk := concrete_cryptography.CreatePublicKeyForTests(t)
	build := createCreateBuilder()
	cr, crErr := build.Create().WithPublicKey(pk).Now()

	if crErr == nil {
		t.Errorf("the returned error was expected to be an instance, nil returned")
	}

	if cr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}

func TestBuildCreate_withID_withoutPublicKey_returnsError(t *testing.T) {

	//variables:
	id := uuid.NewV4()

	//execute:
	build := createCreateBuilder()
	cr, crErr := build.Create().WithID(&id).Now()

	if crErr == nil {
		t.Errorf("the returned error was expected to be an instance, nil returned")
	}

	if cr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}

func TestBuildCreate_withoutID_withoutPublicKey_returnsError(t *testing.T) {

	//execute:
	build := createCreateBuilder()
	cr, crErr := build.Create().Now()

	if crErr == nil {
		t.Errorf("the returned error was expected to be an instance, nil returned")
	}

	if cr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}
