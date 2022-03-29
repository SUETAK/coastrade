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
	SelectTicker() ([]model.Ticker, error)
	CreateCandleTable(model.Candle)
	SelectCandle(candle model.Candle, time time.Time) (*model.Candle, error)
	SelectAllCandle(candle model.Candle, limit int) ([]model.Candle, error)
	InsertCandle(candle *model.Candle)
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

func (c CryptoDB) InsertCandle(candle *model.Candle) {
	tableName := candle.ProductCode + candle.Duration.String()
	_, err := c.db.NewInsert().Model(candle).ModelTableExpr(tableName).Exec(context.Background())
	if err != nil {
		err.Error()
	}
}

func (c CryptoDB) CreateCandleTable(candle model.Candle) {
	tableName := candle.ProductCode + candle.Duration.String()
	_, err := c.db.NewCreateTable().Model(&model.Candle{}).ModelTableExpr(tableName).IfNotExists().Exec(context.Background())
	fmt.Sprintf("create Candle table %s", tableName)
	if err != nil {
		// TODO panic だとアプリ全体が落ちるので、エラーだけ出力する
		panic(err)
	}
}

func (c CryptoDB) SelectCandle(candle model.Candle, time time.Time) (*model.Candle, error) {
	tableName := candle.ProductCode + candle.Duration.String()
	getCandle := model.Candle{}
	if err := c.db.NewSelect().Model(&getCandle).Table(tableName).Where("time = ?", bun.Ident(time.String())).Order("timestamp ASC").Scan(context.Background()); err != nil {
		return nil, err
	}
	return &getCandle, nil
}

func (c CryptoDB) SelectAllCandle(candle model.Candle, limit int) ([]model.Candle, error) {
	candles := make([]model.Candle, 0)
	tableName := candle.ProductCode + candle.Duration.String()
	if err := c.db.NewSelect().Model(&candles).Table(tableName).Order("times DESC").Limit(limit).Scan(context.Background()); err != nil {
		panic(err)
	}
	return candles, nil
}

func (c CryptoDB) SelectTicker() (ticker []model.Ticker, err error) {
	tickers := make([]model.Ticker, 0)
	if err := c.db.NewSelect().Model(&tickers).Order("timestamp ASC").Scan(context.Background()); err != nil {
		panic(err)
	}
	return tickers, nil
}
