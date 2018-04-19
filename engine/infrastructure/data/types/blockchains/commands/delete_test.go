package commands

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateDelete_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Delete)
	obj := CreateDeleteForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
