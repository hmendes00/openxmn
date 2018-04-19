package commands

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateInsert_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Insert)
	obj := CreateInsertForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
