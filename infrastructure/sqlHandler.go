// Package infrastructure domain>repository で定義されたinterface 実装を各所
package infrastructure

import (
	"coastrade/domain/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() CryptoSQL {
	return CryptoDB{db: createSqlConnect()}
}

type CryptoSQL interface {
	CloseDBConnection()
	CreateNewTable() bool
	InsertTicker(ticker *model.Ticker)
}

type CryptoDB struct {
	db *bun.DB
}

//多分Injectパッケージの役割
func createSqlConnect() *bun.DB {
	DBMS := "mysql"
	USER := "suetak"
	PASSWORD := "suetak"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "crypto"

	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME

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

func (c CryptoDB) CloseDBConnection() {
	err := c.db.Close()
	if err != nil {
		panic(err)
	}
}
func (c CryptoDB) CreateNewTable() bool {
	_, err := c.db.NewCreateTable().Model((*model.Ticker)(nil)).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("create_table")
	return true
}

func (c CryptoDB) InsertTicker(ticker *model.Ticker) {
	_, err := c.db.NewInsert().Model(ticker).Exec(context.Background())
	if err != nil {
		// TODO panic だとアプリ全体が落ちるので、エラーだけ出力する
		panic(err)
	}
}
