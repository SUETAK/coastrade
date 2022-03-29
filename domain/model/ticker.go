package model

import (
	"log"
	"time"
)

type Ticker struct {
	ProductCode     string  `json:"product_code"     bun:"product_code"`
	State           string  `json:"state"            bun:"state"`
	Timestamp       string  `json:"timestamp"        bun:"time_stamp"`
	TickID          int     `json:"tick_id"          bun:"tick_id"`
	BestBid         float64 `json:"best_bid"          bun:"best_bid"`        // 板寄せ時売りの成り行き注文最大値
	BestAsk         float64 `json:"best_ask"          bun:"best_ask"`        // 板寄せ時買いの成り行き注文最大値
	BestBidSize     float64 `json:"best_bid_size"     bun:"best_bid_size"`   // 板寄せ時の売りの成行注文量
	BestAskSize     float64 `json:"best_ask_size"     bun:"best_ask_size"`   // 板寄せ時の買いの成行注文量
	TotalBidDepth   float64 `json:"total_bid_depth"   bun:"total_bid_depth"` // 買いの指値注文の数量
	TotalAskDepth   float64 `json:"total_ask_depth"   bun:"total_ask_depth"` // 売りの指値注文の数量
	MarketBidSize   float64 `json:"market_bid_size"   bun:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"   bun:"market_ask_size"`
	Ltp             float64 `json:"ltp"               bun:"ltp"`    // 最終取引価格
	Volume          float64 `json:"volume"            bun:"volume"` // 24時間の取引量
	VolumeByProduct float64 `json:"volume_by_product" bun:"volume_by_product"`
}

func (t Ticker) TruncateDateTime(duration time.Duration) time.Time {
	return t.dateTime().Truncate(duration)
}

func (t Ticker) dateTime() time.Time {
	date, err := time.Parse(time.RFC3339, t.Timestamp)
	if err != nil {
		log.Printf("action=DateTime, err=%s", err.Error())
	}
	return date
}

func (t Ticker) GetMidPrice() float64 {
	return (t.BestBid + t.BestAsk) / 2
}
