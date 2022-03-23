package main

import (
	config "coastrade/configs"
	handler "coastrade/handler/rest"
	"coastrade/infrastructure"
	"coastrade/usecase"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	db := infrastructure.NewDB()
	defer func(db infrastructure.CryptoSQL) {
		db.CloseDBConnection()
	}(db)
	config := config.CreateConfig()
	// bitflyerのapiにリクエストして、レスポンスを受け取る
	fmt.Println(config)

	tickerHandler := NewTicker(db, config)

	router := httprouter.New()
	router.GET("/api/ticker", tickerHandler.Index)
	router.GET("/api/continue/ticker", tickerHandler.ContinueIndex)

	port := ":8087"
	fmt.Println(`Server Start >> http://localhost:%s`, port)
	log.Fatal(http.ListenAndServe(port, router))

}

func NewTicker(db infrastructure.CryptoSQL, config config.Config) handler.TickerHandler {
	return handler.NewTickerHandler(usecase.NewTickerUseCase(db, config))
}
