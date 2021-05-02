package api

import "net/http"

func DoRequest() string {
	return "hello go"
}

// この動作や、値しか許容しない構造体にする
type Client struct {
	apikey string
	secretkey string
	httpClient *http.Client
}

// レスポンスを受け取る構造体
type Ticker struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         int     `json:"best_bid"`
	BestAsk         int     `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     int     `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   int     `json:"total_ask_depth"`
	MarketBidSize   int     `json:"market_bid_size"`
	MarketAskSize   int     `json:"market_ask_size"`
	Ltp             int     `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}