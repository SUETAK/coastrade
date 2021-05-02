package infrastructure

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "suetak"
	PASSWORD := "suetak"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "crypto"

	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 60 {
				fmt.Print("")
				fmt.Print("db接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")
	return db
}
