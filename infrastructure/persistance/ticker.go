package persistance

import (
	"coastrade/api/client"
	"coastrade/domain"
	"coastrade/domain/repository"
)

type tickerPersistance struct {}

func NewTickerPersistance() repository.TickerRepository {
	return &tickerPersistance{}
}

func (ticker tickerPersistance) GetTicker() (*domain.Ticker, error){
	_, err := client.DoRequest("getticker")
	if err != nil {
		return nil, err
	}
	return &domain.Ticker{}, nil
}