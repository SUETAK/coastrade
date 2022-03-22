// Package repository ドメインモデルオブジェクトに関するロジックを定義する
package repository

import "coastrade/domain/model"

type TickerRepository interface {
	GetTicker() (*model.Ticker, error)
}
