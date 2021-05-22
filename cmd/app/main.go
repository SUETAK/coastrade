package main

import (
	config "coastrade/configs"
	handler "coastrade/handler/rest"
	"coastrade/infrastructure/persistance"
	"coastrade/usecase"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// db := persistance.SqlConnect()
	// defer db.Close()
	configUser := config.Config.User
	// bitflyerのapiにリクエストして、レスポンスを受け取る
	fmt.Println(configUser)

	tickerPersistance := persistance.NewTickerPersistance()
	tickerUseCase := usecase.NewTickerUseCase(tickerPersistance)
	tickerHandler := handler.NewTickerHandler(tickerUseCase)

	router := httprouter.New()
	router.GET("/api/ticker", tickerHandler.Index)

	port := ":8086"
	fmt.Println(`Server Start >> http:// localhost:%d`, port)
	log.Fatal(http.ListenAndServe(port, router))

}
