package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/xegcrbq/P2PChat/internal/api"
	"os"
	"os/signal"
	"time"
)

const (
	listenAddr       = "0.0.0.0:8080"
	postgressDb      = "postgres://root:password@db:5432/docker"
	shutdown_timeout = 5
)

func main() {
	fmt.Println(time.Now())
	time.Sleep(time.Second * 8)
	//2022-09-28 20:23:27.4110301
	log := logrus.New()
	//fmt.Println("start")
	// -------------------- Set up database -------------------- //

	dbPool, err := pgxpool.New(context.Background(), postgressDb)
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}
	defer dbPool.Close()
	//fmt.Println("db")
	// -------------------- Set up service -------------------- //

	svc, err := api.NewAPIService(logrus.NewEntry(log), dbPool)
	if err != nil {
		log.Fatalf("error creating service instance: %s", err)
	}
	go svc.Serve(listenAddr)
	//fmt.Println("service")

	// -------------------- Listen for INT signal -------------------- //

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	//fmt.Println("listen")
	<-quit

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*time.Duration(shutdown_timeout),
	)
	defer cancel()

	if err := svc.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("end")
}
