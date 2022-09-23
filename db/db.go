package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/url"
	"os"
)

func ConnectDB() *sqlx.DB {
	q := url.Values{}
	q.Set("sslmode", "disable")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD")),
		Host:     os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT"), // change here
		Path:     os.Getenv("POSTGRES_DB"),
		RawQuery: q.Encode(),
	}
	dbDriver := "postgres"
	db, err := sqlx.Open(dbDriver, u.String())
	if err != nil {
		panic(err.Error())
	}
	return db
}
func ConnectRedis() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return r
}
func ConnectSQLXTest() *sqlx.DB {
	q := url.Values{}
	q.Set("sslmode", "disable")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword("postgres", "1"),
		Host:     "localhost:5432", // change here
		Path:     "postgres",
		RawQuery: q.Encode(),
	}
	dbDriver := "postgres"
	db, err := sqlx.Open(dbDriver, u.String())
	if err != nil {
		panic(err.Error())
	}
	return db
}
