package aggregated

import (
	"testing"

	convert "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateSignedTransactions_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(SignedTransactions)
	obj := CreateSignedTransactionsForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
