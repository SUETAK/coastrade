package model

import "math"

// Balance 自分が保有しているコイン、現金の量と、取引できる量
type Balance struct {
	CurrentCode string  `json:"currency_code"`
	Amount      float64 `json:"amount"`
	Available   float64 `json:"available"`
}

func (b *Balance) GetAvailable(currencyCode string) (availableCurrency float64) {
	if b.CurrentCode == currencyCode {
		return b.Available
	}
	return 0
}

func AdjustSize(size float64) float64 {
	fee := size * 0.0012
	size = size - fee
	return math.Floor(size*10000) / 10000
}
