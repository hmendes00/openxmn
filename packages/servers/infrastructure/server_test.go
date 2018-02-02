package infrastructure

import (
	"net"
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateServer_Success(t *testing.T) {

	//variables:
	protocol := "http"
	port := 80
	ip := net.IPv4(127, 0, 0, 1)

	//execute:
	serv := createServer(protocol, port, ip)
	retIP := serv.GetIP()
	retPort := serv.GetPort()
	retProtocol := serv.GetProtocol()
	retStr := serv.String()

	if !reflect.DeepEqual(ip, retIP) {
		t.Errorf("the returned IP is invalid")
	}

	if !reflect.DeepEqual(port, retPort) {
		t.Errorf("the returned port is invalid")
	}

	if !reflect.DeepEqual(protocol, retProtocol) {
		t.Errorf("the returned protocol is invalid")
	}

	if retStr != "http://127.0.0.1:80" {
		t.Errorf("the returned ip as string (%s) is invalid", retStr)
	}

}

func TestCreateServer_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Server)
	obj := CreateServerForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
