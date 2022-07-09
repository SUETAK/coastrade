package trade

import (
	"coastrade/api/client"
	"coastrade/domain/model"
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
	balances, err := u.client.GetBalance(product)
	if err != nil {
		return nil, err
	}

	var availableSize float64
	for _, balance := range balances {
		availableSize = balance.GetAvailable(product)
	}
	if availableSize == 0 {
		return &infrastructure.ResponseSendChildOrder{ChildOrderAcceptanceID: ""}, nil
	}
	if err != nil {
		return nil, err
	}
	if decidedPosition == "buy" {
		buyOrder := &infrastructure.Order{
			ProductCode:     product,
			ChildOrderType:  "MARKET",
			Side:            "SELL",
			Size:            AdjustSize(),
			MinuteToExpires: 100,
			TimeInForce:     "GTC",
		}
		resp, err = u.client.SendOrder(buyOrder, product)
		if err != nil {
			return nil, err
		}
	}
	if decidedPosition == "sell" {
		sellOrder := &infrastructure.Order{}
		resp, err = u.client.SendOrder(sellOrder, "ETH")
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func AdjustSize(size float64) float64 {
	fee := size * 0.0012
	size = size - fee
	return math.Floor(size*10000) / 10000
}
