package rest

import (
	"coastrade/usecase"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TickerHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type tickerHandler struct {
	tickerUseCase usecase.TickerUseCase
}

func NewTickerHandler(tu usecase.TickerUseCase) TickerHandler {
	return &tickerHandler{
		tickerUseCase: tu,
	}
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
