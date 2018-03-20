package users

import (
	"testing"

	convert "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateSignatures_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Signatures)
	obj := CreateSignaturesForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
