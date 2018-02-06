package custom

import (
	"reflect"
	"testing"
)

func TestBuildCustom_withCreate_Success(t *testing.T) {

	//variables:
	cr := CreateCreateForTests(t)

	//execute:
	cu, cuErr := createCustomBuilder().Create().WithCreate(cr).Now()

	if cuErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", cuErr.Error())
	}

	if cu == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

	retCreate := cu.GetCreate()

	if !reflect.DeepEqual(cr, retCreate) {
		t.Errorf("the returned create is invalid.")
	}

}

func TestBuildCustom_withoutTrs_Success(t *testing.T) {

	//execute:
	cu, cuErr := createCustomBuilder().Create().Now()

	if cuErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if cu != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
