package main

import (
	"coastrade/api/client"
	config "coastrade/configs"
	"coastrade/infrastructure/persistance"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := persistance.SqlConnect()
	defer db.Close()
	configUser := config.Config.User
	// bitflyerのapiにリクエストして、レスポンスを受け取る
	fmt.Println(configUser)
	apiClient := client.New(config.Config.ApiKey,
		 config.Config.ApiSecret)
	apiClient.DoRequest("ticker")

}
