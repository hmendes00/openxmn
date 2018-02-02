package infrastructure

import (
	"net"
	"net/url"
	"reflect"
	"testing"
)

func TestBuildServer_withURI_Success(t *testing.T) {

	//variables:
	ur, _ := url.Parse("http://127.0.0.1:80")
	protocol := "http"
	port := 80
	ip := net.IPv4(127, 0, 0, 1)

	//execute:
	build := createServerBuilder()
	serv, servErr := build.Create().WithURL(ur).Now()
	if servErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", servErr.Error())
	}

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

func TestBuildServer_withoutURI_returnsError(t *testing.T) {

	//execute:
	build := createServerBuilder()
	serv, servErr := build.Create().Now()

	if servErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if serv != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildServer_withInvalidPort_returnsError(t *testing.T) {

	ur, _ := url.Parse("http://[2001:0:53aa:64c:104c:2c10:2bef:4f7b]")

	//execute:
	build := createServerBuilder()
	serv, servErr := build.Create().WithURL(ur).Now()

	if servErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if serv != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildServer_withNonIPHostname_returnsError(t *testing.T) {

	ur, _ := url.Parse("http://not-an-ip.com:80")

	//execute:
	build := createServerBuilder()
	serv, servErr := build.Create().WithURL(ur).Now()

	if servErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if serv != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
