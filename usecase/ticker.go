// Package usecase domain を使ったロジックを提供するパッケージ
package usecase

import (
	"coastrade/domain/model"
	"coastrade/infrastructure/persistence"
)

type TickerUseCase interface {
	GetTicker() (*model.Ticker, error)
}

type tickerUseCase struct {
	tickerPersistence persistence.TickerPersistence
}

func (tu *tickerUseCase) GetTicker() (ticker *model.Ticker, err error) {
	ticker, err = tu.tickerPersistence.GetTicker()
	if err != nil {
		return nil, err
	}
	return ticker, nil
}

func NewTickerUseCase(tp persistence.TickerPersistence) TickerUseCase {
	return &tickerUseCase{
		tickerPersistence: tp,
	}
}
