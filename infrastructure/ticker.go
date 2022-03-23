// Package infrastructure 外部への疎通などを扱う層。ticker関連の操作はこのファイルで行う
package infrastructure

import (
	"coastrade/api/client"
	config "coastrade/configs"
	"errors"
	"fmt"

	"coastrade/domain/model"
	"encoding/json"
	"log"
)

func NewTickerInfra() TickerInfra {
	return &ticker{}
}

type TickerInfra interface {
	GetTicker() (*model.Ticker, error)
}

type ticker struct{}

func (tp *ticker) GetTicker() (*model.Ticker, error) {
	apiClient := client.New(config.Config.ApiKey, config.Config.ApiSecret)
	response, err := apiClient.DoRequest("ticker", "GET")
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	responseBytes := []byte(response)

	fmt.Printf("%s", response)

	var ticker model.Ticker
	err = json.Unmarshal(responseBytes, &ticker)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error")
	}
	return &ticker, nil
}
