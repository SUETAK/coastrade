package rest

import (
	"coastrade/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func NewTickerHandler(tu usecase.TickerUseCase) TickerHandler {
	return &tickerHandler{
		tickerUseCase: tu,
	}
}

type TickerHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
	ContinueIndex(http.ResponseWriter, *http.Request, httprouter.Params)
}

type tickerHandler struct {
	tickerUseCase usecase.TickerUseCase
}

func (th tickerHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	ticker, err := th.tickerUseCase.GetTicker()

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
}

func (th tickerHandler) ContinueIndex(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	go func() {
		//クライアントにレスポンスを返却
		w.Header().Set("Content-Type", "application/json")
		for {
			fmt.Printf("now -> %v\n", time.Now())
			_, err := th.tickerUseCase.GetTicker()
			if err != nil {
				http.Error(w, "Internal Sever Error", 500)
				return
			}

			time.Sleep(time.Second * 3)
		}
	}()
}
