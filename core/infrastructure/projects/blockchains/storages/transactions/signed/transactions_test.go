package signed

import (
	"testing"

	convert "github.com/XMNBlockchain/exmachina-network/core/infrastructure/tests/jsonify/helpers"
)

func TestCreateTransactions_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transactions)
	obj := CreateTransactionsForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
