package signed

import (
	"testing"

	convert "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateAtomicTransactions_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(AtomicTransactions)
	obj := CreateAtomicTransactionsForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
