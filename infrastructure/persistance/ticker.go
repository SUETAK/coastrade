package persistance

import (
	"coastrade/api/client"
	config "coastrade/configs"
	"coastrade/domain"
	"coastrade/domain/repository"
)

type tickerPersistance struct {}

func NewTickerPersistance() repository.TickerRepository {
	return &tickerPersistance{}
}

func (ticker tickerPersistance) GetTicker() (*domain.Ticker, error){
	apiClient := client.New(config.Config.ApiKey, config.Config.ApiSecret)
	_, err := apiClient.DoRequest("getticker")
	if err != nil {
		return nil, err
	}
	return &domain.Ticker{}, nil
}