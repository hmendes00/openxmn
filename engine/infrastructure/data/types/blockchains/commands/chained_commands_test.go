package commands

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateChainedCommands_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(ChainedCommands)
	obj := CreateChainedCommandsForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
