package chained

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateBlock_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Block)
	obj := CreateBlockForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
