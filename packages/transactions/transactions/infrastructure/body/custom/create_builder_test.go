package custom

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestBuildCreate_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	insID := uuid.NewV4()
	ins := InstanceForTests{
		ID:          &insID,
		Name:        "some name",
		Description: "this is some custom description",
	}
	js, jsErr := json.Marshal(ins)
	if jsErr != nil {
		t.Errorf("there was an error while converting an instance to json: %s", jsErr)
	}

	//execute:
	cr, crErr := createCreateBuilder().Create().Create().WithID(&id).WithInstance(ins).Now()

	if crErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", crErr.Error())
	}

	if cr == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

	retID := cr.GetID()
	retJS := cr.GetJSON()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID is invalid.  Returned: %s, Expected: %s", retID.String(), id.String())
	}

	if !reflect.DeepEqual(js, retJS) {
		t.Errorf("the returned JSON is invalid")
	}

	newIns := new(InstanceForTests)
	newErr := json.Unmarshal(retJS, newIns)
	if newErr != nil {
		t.Errorf("there was an error while converting json to an instance: %s", newErr)
	}

	if !reflect.DeepEqual(newIns, &ins) {
		t.Errorf("the returned JSON is invalid because it could not re-create the instance")
	}

}

func TestBuildCreate_withoutID_returnsError(t *testing.T) {

	//variables:
	insID := uuid.NewV4()
	ins := InstanceForTests{
		ID:          &insID,
		Name:        "some name",
		Description: "this is some custom description",
	}

	//execute:
	cr, crErr := createCreateBuilder().Create().Create().WithInstance(ins).Now()

	if crErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if cr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildCreate_withoutInstance_returnsError(t *testing.T) {

	//variables:
	id := uuid.NewV4()

	//execute:
	cr, crErr := createCreateBuilder().Create().Create().WithID(&id).Now()

	if crErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if cr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildCreate_withID_withInvalidInstance_returnsError(t *testing.T) {

	//variables:
	id := uuid.NewV4()

	//execute:
	cr, crErr := createCreateBuilder().Create().Create().WithID(&id).WithInstance(math.Inf(1)).Now()

	if crErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if cr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
