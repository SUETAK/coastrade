package trade

import (
	"coastrade/api/client"
	"coastrade/domain/model"
	"coastrade/infrastructure"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type APIClientMock struct {
	mock.Mock
}

func (m *APIClientMock) SendOrder(order *infrastructure.Order, product string) (*infrastructure.ResponseSendChildOrder, error) {
	args := m.Called(order, product)
	return args.Get(0).(*infrastructure.ResponseSendChildOrder), args.Error(1)
}

func (m APIClientMock) ListOrder(query map[string]string, product string) ([]infrastructure.Order, error) {
	args := m.Called(query, product)
	return args.Get(0).([]infrastructure.Order), args.Error(1)
}

type DecideMock struct {
	mock.Mock
}

func (m DecideMock) DecidePosition(trade *criteria, value float64) (string, error) {
	args := m.Called(trade, value)
	return args.Get(0).(string), args.Error(1)
}

type TickerMock struct {
	mock.Mock
}

func (m TickerMock) GetTicker(product string) (*model.Ticker, error) {
	args := m.Called(product)
	return args.Get(0).(*model.Ticker), args.Error(1)
}

func Test_tradeUsecase_DoTrading(t *testing.T) {
	type fields struct {
		ticker   infrastructure.TickerInfra
		position Decide
		client   client.APIClient
	}
	tickerMock := new(TickerMock)
	positionMock := new(DecideMock)
	clientMock := new(APIClientMock)

	tests := []struct {
		name    string
		fields  fields
		want    *infrastructure.ResponseSendChildOrder
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			"てすと",
			fields{
				ticker:   tickerMock,
				position: positionMock,
				client:   clientMock,
			},
			&infrastructure.ResponseSendChildOrder{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tradeUsecase{
				ticker:   tt.fields.ticker,
				position: tt.fields.position,
				client:   tt.fields.client,
			}
			got, err := u.DoTrading()
			if !tt.wantErr(t, err, fmt.Sprintf("DoTrading()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "DoTrading()")
		})
	}
}
