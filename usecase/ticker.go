// Package usecase domain を使ったロジックを提供するパッケージ
package usecase

import (
	"coastrade/domain/model"
	"coastrade/infrastructure"
)

func NewTickerUseCase(db infrastructure.CryptoSQL) TickerUseCase {
	return &tickerUseCase{
		tickerInfra: infrastructure.NewTickerInfra(),
		cryptoSQL:   db,
	}
}

type TickerUseCase interface {
	GetTicker() (*model.Ticker, error)
}

type tickerUseCase struct {
	tickerInfra infrastructure.TickerInfra
	cryptoSQL   infrastructure.CryptoSQL
}

func (tu *tickerUseCase) GetTicker() (ticker *model.Ticker, err error) {
	ticker, err = tu.tickerInfra.GetTicker()
	if err != nil {
		return nil, err
	}
	tu.cryptoSQL.InsertTicker(ticker)
	return ticker, nil
}
