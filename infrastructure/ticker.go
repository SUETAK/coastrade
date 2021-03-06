// Package infrastructure 外部への疎通などを扱う層。ticker関連の操作はこのファイルで行う
package infrastructure

import (
	"coastrade/api/client"
	config "coastrade/configs"
	"coastrade/domain/model"
	"encoding/json"
	"errors"
	"log"
)

func NewTickerInfra(config config.Config) TickerInfra {
	return &ticker{
		config: config,
	}
}

type TickerInfra interface {
	GetTicker(product string) (*model.Ticker, error)
}

type ticker struct {
	config config.Config
}

func (tp *ticker) GetTicker(product string) (*model.Ticker, error) {
	apiClient := client.New(tp.config.ApiKey, tp.config.ApiSecret)
	response, err := apiClient.DoRequest("ticker", "GET", product, nil, nil)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}

	var responseTicker model.Ticker
	err = json.Unmarshal(response, &responseTicker)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error")
	}
	return &responseTicker, nil
}

// TODO ロジックを実装する
func (tp *ticker) Buy() {

}

func (tp *ticker) Sell() {

}
