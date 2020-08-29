package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx"
)

var dbConn *pgx.ConnPool
var redisConn *redis.Client
var data map[string]string

func init() {
	var err error

	pgxConf := &pgx.ConnConfig{
		Port:     5432,
		Host:     "localhost",
		User:     "omama",
		Password: "omama",
		Database: "omama",
	}

	if dbConn, err = pgx.NewConnPool(
		pgx.ConnPoolConfig{
			ConnConfig:     *pgxConf,
			MaxConnections: 5,
		},
	); err != nil {
		panic(err)
	}

	redisConn = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

}

func main() {
	fmt.Println(dbConn, redisConn)
}

func setToCache(datas data) {

}

func loadFromCache() (datas data) {
	return
}

func loadFromDB() (datas data) {
	return
}
