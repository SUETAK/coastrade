package persistance

// APIの技術的関心ごとを扱う層。ticker関連の操作はここで行う
import (
	"coastrade/api/client"
	config "coastrade/configs"
	"errors"
	"fmt"

	"coastrade/domain/model"
	"coastrade/domain/repository"
	"encoding/json"
	"log"
)

type TickerPersistance struct{}

// 返り値をインターフェース型であるTickerRepositoryに指定。
// TickerRepositoryに依存していることになる
func NewTickerPersistance() repository.TickerRepository {
	return &TickerPersistance{}
}

func (tp TickerPersistance) GetTicker() (*model.Ticker, error) {
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
