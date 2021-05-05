package repository

import "coastrade/domain"

type TickerRepository interface {
	GetTicker() (*domain.Ticker, error) 
}