package commands

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateCommand_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Command)
	obj := CreateCommandForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
