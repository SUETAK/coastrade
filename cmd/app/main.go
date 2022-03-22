package main

import (
	config "coastrade/configs"
	handler "coastrade/handler/rest"
	"coastrade/infrastructure/persistence"
	"coastrade/usecase"
	"fmt"
	"github.com/uptrace/bun"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := persistence.SqlConnect()
	defer func(db *bun.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	configUser := config.Config.User
	// bitflyerのapiにリクエストして、レスポンスを受け取る
	fmt.Println(configUser)

	tickerHandler := NewTicker()

	router := httprouter.New()
	router.GET("/api/ticker", tickerHandler.Index)

	port := ":8087"
	fmt.Println(`Server Start >> http://localhost:%s`, port)
	log.Fatal(http.ListenAndServe(port, router))

}

func NewTicker() handler.TickerHandler {
	tickerPersistence := persistence.NewTickerPersistence()
	tickerUseCase := usecase.NewTickerUseCase(*tickerPersistence)
	return handler.NewTickerHandler(tickerUseCase)
}
