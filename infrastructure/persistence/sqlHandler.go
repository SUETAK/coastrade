// Package persistence domain>repository で定義されたinterface 実装を各所
package persistence

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//多分Injectパッケージの役割
func SqlConnect() *bun.DB {
	//TODO Bun に移行する
	DBMS := "mysql"
	USER := "suetak"
	PASSWORD := "suetak"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "crypto"

	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	sqldb, err := sql.Open(DBMS, CONNECT)
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
			sqldb, err = sql.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")
	return bun.NewDB(sqldb, mysqldialect.New())
}

func CreateNewTable() bool {
	// TODO 対象のテーブルがない場合は新しいテーブルを作る
	return true
}
