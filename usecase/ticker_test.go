package usecase

import (
	"coastrade/domain/model"
	"coastrade/infrastructure"
	"reflect"
	"testing"
)

func createTickerMock() infrastructure.TickerInfra {
	return &tickerMock{}
}

type tickerMock struct {
}

func (receiver tickerMock) GetTicker() (*model.Ticker, error) {
	return nil, nil
}

func createCryptoSQLMock() infrastructure.CryptoSQL {
	return &cryptoSQLMock{}
}

type cryptoSQLMock struct {
}

func (receiver cryptoSQLMock) CloseDBConnection() {

}

func (receiver cryptoSQLMock) CreateNewTable() bool {
	return true
}

func (receiver cryptoSQLMock) InsertTicker(ticker *model.Ticker) {
}

func Test_tickerUseCase_GetTicker(t *testing.T) {
	// TODO config ファイルに依存しないようにusecase を書き換える
	type fields struct {
		tickerInfra infrastructure.TickerInfra
		cryptoSQL   infrastructure.CryptoSQL
	}
	tests := []struct {
		name       string
		fields     fields
		wantTicker *model.Ticker
		wantErr    bool
	}{
		{
			name: "call_insert",
			fields: fields{
				tickerInfra: createTickerMock(),
				cryptoSQL:   createCryptoSQLMock(),
			},
			wantTicker: &model.Ticker{
				ProductCode:     "",
				State:           "",
				Timestamp:       "",
				TickID:          0,
				BestBid:         0,
				BestAsk:         0,
				BestBidSize:     0,
				BestAskSize:     0,
				TotalBidDepth:   0,
				TotalAskDepth:   0,
				MarketBidSize:   0,
				MarketAskSize:   0,
				Ltp:             0,
				Volume:          0,
				VolumeByProduct: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tu := &tickerUseCase{
				tickerInfra: tt.fields.tickerInfra,
				cryptoSQL:   tt.fields.cryptoSQL,
			}
			gotTicker, err := tu.GetTicker()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTicker, tt.wantTicker) {
				t.Errorf("GetTicker() gotTicker = %v, want %v", gotTicker, tt.wantTicker)
			}
		})
	}
}
