// Package usecase domain を使ったロジックを提供するパッケージ
package usecase

import (
	config "coastrade/configs"
	"coastrade/domain/model"
	"coastrade/infrastructure"
	"time"
)

func NewTickerUseCase(db infrastructure.CryptoSQL, config config.Config) TickerUseCase {
	return &tickerUseCase{
		tickerInfra: infrastructure.NewTickerInfra(config),
		cryptoSQL:   db,
	}
}

type TickerUseCase interface {
	GetTicker(string, time.Duration) (*model.Ticker, error)
}

type tickerUseCase struct {
	tickerInfra infrastructure.TickerInfra
	cryptoSQL   infrastructure.CryptoSQL
}

func (tu *tickerUseCase) GetTicker(product string, duration time.Duration) (ticker *model.Ticker, err error) {
	ticker, err = tu.tickerInfra.GetTicker(product)
	if err != nil {
		return nil, err
	}
	tu.cryptoSQL.InsertTicker(ticker)
	midPrice := ticker.GetMidPrice()
	truncateDateTime := ticker.TruncateDateTime(duration)
	candle := model.Candle{
		ProductCode: ticker.ProductCode,
		Duration:    duration,
		Time:        truncateDateTime,
		Open:        midPrice,
		Close:       midPrice,
		High:        midPrice,
		Low:         midPrice,
		Volume:      ticker.Volume,
	}
	tu.cryptoSQL.CreateCandleTable(candle)
	if err != nil {
		err.Error()
	}
	tu.cryptoSQL.InsertCandle(candle)
	return ticker, nil
}

func (tu *tickerUseCase) Buy() {

}

func (tu tickerUseCase) Sell() {

}
