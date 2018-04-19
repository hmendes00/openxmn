package commands

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateError_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Error)
	obj := CreateErrorForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
