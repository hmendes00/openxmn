package commands

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateUpdate_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Update)
	obj := CreateUpdateForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
