package trade

import (
	"coastrade/api/client"
	"coastrade/infrastructure"
)

type Trade interface {
	DoTrading() (*infrastructure.ResponseSendChildOrder, error)
}

func NewTradeUsecase(ticker infrastructure.TickerInfra, position Decide, client client.APIClient) Trade {
	return tradeUsecase{
		ticker:   ticker,
		position: position,
		client:   client,
	}
}

type tradeUsecase struct {
	ticker   infrastructure.TickerInfra
	position Decide
	client   client.APIClient
}

func (u tradeUsecase) DoTrading() (*infrastructure.ResponseSendChildOrder, error) {
	baseCriteria := NewCriteria(0, 0)

	value, err := u.ticker.GetTicker("ETH")
	if err != nil {

	}
	decidedPosition, err := u.position.DecidePosition(baseCriteria, value.BestAskSize)
	if err != nil {

	}
	var resp *infrastructure.ResponseSendChildOrder
	if decidedPosition == "buy" {
		buyOrder := &infrastructure.Order{}
		resp, err = u.client.SendOrder(buyOrder, "ETH")
		if err != nil {

		}
	}
	if decidedPosition == "sell" {
		sellOrder := &infrastructure.Order{}
		resp, err = u.client.SendOrder(sellOrder, "ETH")
		if err != nil {

		}
	}
	return resp, nil
}
