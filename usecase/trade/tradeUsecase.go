package trade

import (
	"coastrade/api/client"
	"coastrade/infrastructure"
)

type Trade interface {
	DoTrading(product string, criteria *criteria) (*infrastructure.ResponseSendChildOrder, error)
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

func (u tradeUsecase) DoTrading(product string, criteria *criteria) (*infrastructure.ResponseSendChildOrder, error) {
	value, err := u.ticker.GetTicker(product)
	if err != nil {
		return nil, err
	}
	decidedPosition, err := u.position.DecidePosition(criteria, value.BestAskSize)
	if err != nil {
		return nil, err
	}
	var resp *infrastructure.ResponseSendChildOrder
	if decidedPosition == "buy" {
		buyOrder := &infrastructure.Order{}
		resp, err = u.client.SendOrder(buyOrder, product)
		if err != nil {
			return nil, err
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
