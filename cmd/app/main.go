package main

import (
	config "coastrade/configs"
	handler "coastrade/handler/rest"
	"coastrade/infrastructure"
	"coastrade/usecase"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/uptrace/bun"
	"log"
	"net/http"
)

func main() {
	db := infrastructure.SqlConnect()
	defer func(db *bun.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	infrastructure.CreateNewTable(db)
	configUser := config.Config.User
	// bitflyerのapiにリクエストして、レスポンスを受け取る
	fmt.Println(configUser)

	tickerHandler := NewTicker()

	router := httprouter.New()
	router.GET("/api/ticker", tickerHandler.Index)
	router.GET("/api/continue/ticker", tickerHandler.ContinueIndex)

	port := ":8087"
	fmt.Println(`Server Start >> http://localhost:%s`, port)
	log.Fatal(http.ListenAndServe(port, router))

}

func NewTicker() handler.TickerHandler {
	return handler.NewTickerHandler(usecase.NewTickerUseCase())
}
