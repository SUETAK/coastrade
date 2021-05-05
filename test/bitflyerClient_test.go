package test

import (
	client "coastrade/api/client"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	testBaseurl, _ := url.ParseRequestURI("http://example.org")
	testClient := client.New("test", "testKey", testBaseurl)
	api, error := client.NewClient("test", "testKey", "http://example.org")
	fmt.Println(error)
	assert.Nil(t, error)
	assert.Equal(t, api, testClient)
}
