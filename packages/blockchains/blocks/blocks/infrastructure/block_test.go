package infrastructure

import (
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateBlock_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Block)
	obj := CreateBlockForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
