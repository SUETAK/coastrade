package rest

import (
	"coastrade/usecase"
	"coastrade/usecase/trade"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func NewTradingHandler(tu usecase.TickerUseCase, tradingUsecase trade.Trade) TradingHandler {
	return &tradingHandler{
		tickerUseCase:  tu,
		tradingUsecase: tradingUsecase,
	}
}

type TradingHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
	ContinueIndex(http.ResponseWriter, *http.Request, httprouter.Params)
}

type tradingHandler struct {
	tickerUseCase  usecase.TickerUseCase
	tradingUsecase trade.Trade
}

func (th tradingHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	duration, err := time.ParseDuration(r.FormValue("duration"))
	if err != nil {
		err.Error()
	}
	product := r.FormValue("product")
	ticker, err := th.tickerUseCase.GetTicker(product, duration)
	if err != nil {
		http.Error(w, "Internal Sever Error", 500)
		return
	}

	//クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(ticker); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	res, err := json.Marshal(ticker)
	if err != nil {
		panic(err)
	}
	w.Write(res)

}

func (th tradingHandler) ContinueIndex(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	go func() {
		//クライアントにレスポンスを返却
		w.Header().Set("Content-Type", "application/json")
		for {
			fmt.Printf("now -> %v\n", time.Now())
			duration, err := time.ParseDuration(r.FormValue("duration"))
			if err != nil {
				err.Error()
			}
			product := r.FormValue("product")
			_, err = th.tickerUseCase.GetTicker(product, duration)
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				return
			}
			time.Sleep(time.Second)
		}
	}()
}

func (th tradingHandler) ObserveValue() {
	go func() {
		criteria := trade.NewCriteria(0, 0)
		_, err := th.tradingUsecase.DoTrading("ETH", criteria)
		if err != nil {

		}
		time.Sleep(time.Second)
	}()
}
