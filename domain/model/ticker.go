package model

type Ticker struct {
	ProductCode     string  `json:"product_code"     bun:"product_code"`
	State           string  `json:"state"            bun:"state"`
	Timestamp       string  `json:"timestamp"        bun:"time_stamp"`
	TickID          int     `json:"tick_id"          bun:"tick_id"`
	BestBid         float64 `json:"best_bid"          bun:"best_bid"`
	BestAsk         float64 `json:"best_ask"          bun:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"     bun:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"     bun:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"   bun:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"   bun:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"   bun:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"   bun:"market_ask_size"`
	Ltp             float64 `json:"ltp"               bun:"ltp"`
	Volume          float64 `json:"volume"            bun:"volume"`
	VolumeByProduct float64 `json:"volume_by_product" bun:"volume_by_product"`
}
