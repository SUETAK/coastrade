package test

import (
	"testing"
	"coastrade/api"
)

func TestDoRequest(t *testing.T) {
	if api.DoRequest() != "hello go!" {
		t.Fatal("doReqsuest should be hello go")
	}
}