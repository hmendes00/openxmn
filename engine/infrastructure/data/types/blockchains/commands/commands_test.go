package commands

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateCommands_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Commands)
	obj := CreateCommandsForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
