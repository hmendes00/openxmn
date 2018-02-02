package helpers

import (
	"encoding/json"
	"reflect"
	"testing"
)

// ConvertToJSON converts an instance to JSON, then converts it back to instance.  Will fail if an error occurs
func ConvertToJSON(t *testing.T, v interface{}, empty interface{}) {
	if v == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

	js, jsErr := json.Marshal(v)
	if jsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", jsErr.Error())
	}

	newErr := json.Unmarshal(js, empty)
	if newErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", newErr.Error())
	}

	if !reflect.DeepEqual(v, empty) {
		t.Errorf("the json conversion (back and forth) did not succeed.  \n Expected: %v, \n Returned: %v", v, empty)
	}
}
