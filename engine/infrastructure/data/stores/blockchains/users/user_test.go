package users

import (
	"testing"

	convert "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateUser_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(User)
	obj := CreateUserForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
