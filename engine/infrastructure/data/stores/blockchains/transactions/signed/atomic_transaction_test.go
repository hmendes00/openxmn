package signed

import (
	"testing"

	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateAtomicTransaction_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(AtomicTransaction)
	obj := CreateAtomicTransactionForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
