package infrastructure

import (
	"go.uber.org/zap/zapcore"
)

type Order struct {
	ID                     int     `json:"id"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
	ProductCode            string  `json:"product_code"`
	ChildOrderType         string  `json:"child_order_type"`
	Side                   string  `json:"side"`
	Price                  float64 `json:"price"`
	Size                   float64 `json:"size"`
	MinuteToExpires        int     `json:"minute_to_expire"`
	TimeInForce            string  `json:"time_in_force"`
	Status                 string  `json:"status"`
	ErrorMessage           string  `json:"error_message"`
	AveragePrice           float64 `json:"average_price"`
	ChildOrderState        string  `json:"child_order_state"`
	ExpireDate             string  `json:"expire_date"`
	ChildOrderDate         string  `json:"child_order_date"`
	OutstandingSize        float64 `json:"outstanding_size"`
	CancelSize             float64 `json:"cancel_size"`
	ExecutedSize           float64 `json:"executed_size"`
	TotalCommission        float64 `json:"total_commission"`
	Count                  int     `json:"count"`
	Before                 int     `json:"before"`
	After                  int     `json:"after"`
}

func (o Order) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("ID", o.ID)
	enc.AddInt("MinuteToExpires", o.MinuteToExpires)
	enc.AddInt("Count", o.Count)
	enc.AddInt("Before", o.Before)
	enc.AddInt("After", o.After)
	enc.AddString("ChildOrderAcceptanceID", o.ChildOrderState)
	enc.AddString("ProductCode", o.ProductCode)
	enc.AddString("ChildOrderType", o.ChildOrderType)
	enc.AddString("Side", o.Side)
	enc.AddString("TimeInForce", o.TimeInForce)
	enc.AddString("ErrorMessage", o.ErrorMessage)
	enc.AddString("ChildOrderState", o.ChildOrderState)
	enc.AddString("ExpireDate", o.ExpireDate)
	enc.AddString("ChildOrderDate", o.ChildOrderDate)
	enc.AddFloat64("Side", o.Price)
	enc.AddFloat64("Price", o.Price)
	enc.AddFloat64("AveragePrice", o.AveragePrice)
	enc.AddFloat64("OutstandingSize", o.OutstandingSize)
	enc.AddFloat64("CancelSize", o.CancelSize)
	enc.AddFloat64("ExecutedSize", o.ExecutedSize)
	enc.AddFloat64("TotalCommission", o.TotalCommission)
	return nil
}

type ResponseSendChildOrder struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}
