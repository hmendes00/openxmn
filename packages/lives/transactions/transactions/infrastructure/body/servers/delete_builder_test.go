package servers

import (
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestBuildDelete_withID_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()

	//execute:
	build := createDeleteBuilder()
	del, delErr := build.Create().WithID(&id).Now()

	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, instance returned: %s", delErr.Error())
	}

	retID := del.GetID()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

}

func TestBuildDelete_withoutID_Success(t *testing.T) {

	//execute:
	build := createDeleteBuilder()
	del, delErr := build.Create().Now()

	if delErr == nil {
		t.Errorf("the returned error was expected to be an instance, nil returned")
	}

	if del != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
