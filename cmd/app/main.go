package main

import (
	"coastrade/configs"
	"coastrade/cmd/app/infrastructure"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := infrastructure.SqlConnect()
	defer db.Close()
	configUser := config.Config.User
	// bitflyerのapiにリクエストして、レスポンスを受け取る
	fmt.Println(configUser)
}


