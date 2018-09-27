package hello_package

import "testing"

func TestHelloWorld(t *testing.T) {
	hello := HelloWorld("gotest")
	if hello != "Hi, gotest" {
		t.Errorf("Testing fail")
	}

	hello = HelloWorld("gotest@")
	if hello != "Hi, gotest" {
		t.Errorf("Testing fail")
	}
}
