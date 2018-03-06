package infrastructure

import (
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateTransactions_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transactions)
	obj := CreateTransactionsForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
