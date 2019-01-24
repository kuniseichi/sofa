package sofa

import (
	"os"
	"testing"
)

func fixtureTest(t *testing.T) {
	src, err := os.Stat("tests/fixtures")
	if os.IsNotExist(err) || src.Mode().IsRegular() {
		t.Skip("tests/fixtures does not exist or is not a directory")
	}
}

func TestClientCertAuthenticatorPassword(t *testing.T) {
	fixtureTest(t)

	auth, err := ClientCertAuthenticatorPassword(
		"tests/fixtures/password.crt",
		"tests/fixtures/password.crt",
		"tests/fixtures/ca.crt",
		"Pa55word")
	if err != nil {
		t.Fatal(err)
	}

	_, err = auth.Client()
	if err != nil {
		t.Fatal(err)
	}
}
