package users

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateSignature_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Signature)
	obj := CreateSignatureForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
