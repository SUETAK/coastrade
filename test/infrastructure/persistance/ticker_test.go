package persistance_test

import (
	"coastrade/infrastructure/persistance"
	"testing"

	"github.com/stretchr/testify/mock"
)

type tickerMock struct {
	mock.Mock
}


func TestGetTicker(t *testing.T) {
	teickerPersistanceMock := new(tickerMock)
	
}