package api

import (
	"errors"
	"log"
	"net/http"
	"net/url"
)

func doRequest() string {
	return "hello go"
}

// この動作や、値しか許容しない構造体にする
type Client struct {
	apikey, secretkey string
	httpClient        *http.Client
	baseUrl           *url.URL
	log               *log.Logger
}

func New(apikey, secretkey string, baseUrl *url.URL) *Client {
	return &Client{
		apikey:     apikey,
		secretkey:  secretkey,
		httpClient: &http.Client{},
		baseUrl:    baseUrl,
		log:        &log.Logger{},
	}
}

func NewClient(apikey, secretkey, baseUrlstr string) (*Client, error) {
	baseurl, err := url.ParseRequestURI(baseUrlstr)
	if err != nil {
		return nil, err
	}
	if len(apikey) == 0 {
		return nil, errors.New("apikey is empty")
	}
	return &Client{
		apikey:     apikey,
		secretkey:  secretkey,
		httpClient: &http.Client{},
		baseUrl:    baseurl,
		log:        &log.Logger{},
	}, nil
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
