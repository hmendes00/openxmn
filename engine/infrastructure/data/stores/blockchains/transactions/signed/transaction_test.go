package signed

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateTransaction_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transaction)
	obj := CreateTransactionForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
