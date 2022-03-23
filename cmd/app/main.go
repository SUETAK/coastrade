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
	configUser := config.Config.User
	// bitflyerのapiにリクエストして、レスポンスを受け取る
	fmt.Println(configUser)

	tickerHandler := NewTicker(db)

	router := httprouter.New()
	router.GET("/api/ticker", tickerHandler.Index)
	router.GET("/api/continue/ticker", tickerHandler.ContinueIndex)

	port := ":8087"
	fmt.Println(`Server Start >> http://localhost:%s`, port)
	log.Fatal(http.ListenAndServe(port, router))

}

func NewTicker(db infrastructure.CryptoSQL) handler.TickerHandler {
	return handler.NewTickerHandler(usecase.NewTickerUseCase(db))
}
