package blocks

import (
	"testing"

	convert "github.com/XMNBlockchain/exmachina-network/core/infrastructure/tests/jsonify/helpers"
)

func TestCreateSignedBlock_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(SignedBlock)
	obj := CreateSignedBlockForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
