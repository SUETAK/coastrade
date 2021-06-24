package usecase

import (
	"coastrade/domain/model"
	"coastrade/domain/repository"
)

type TickerUseCase interface {
	GetTicker() (*model.Ticker, error)
}

type tickerUseCase struct {
	tickerRepository repository.TickerRepository
}

func NewTickerUseCase(tr repository.TickerRepository) TickerUseCase {
	return &tickerUseCase{
		tickerRepository: tr,
	}
}

func (tu tickerUseCase) GetTicker() (ticker *model.Ticker, err error) {
	ticker, err = tu.tickerRepository.GetTicker()
	if err != nil {
		return nil, err
	}
	return ticker, nil
}
