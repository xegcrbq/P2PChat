package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/xegcrbq/P2PChat/internal/api"
	"os"
	"os/signal"
	"time"
)

const (
	listenAddr       = "0.0.0.0:8080"
	postgressDb      = "postgres://root:password@localhost:5432/docker"
	shutdown_timeout = 5
	X_Token          = "xuw9xn7znrz4658f862quecb1p8n1s32vhpo35m61yzrofjepnqk0i2tlum3vhqr"
)

func main() {
	log := logrus.New()
	// -------------------- Set up database -------------------- //

	dbPool, err := pgxpool.New(context.Background(), postgressDb)
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err)
	}
	defer dbPool.Close()

	// -------------------- Set up service -------------------- //

	svc, err := api.NewAPIService(logrus.NewEntry(log), dbPool)
	if err != nil {
		log.Fatalf("error creating service instance: %s", err)
	}
	go svc.Serve(listenAddr)

	// -------------------- Listen for INT signal -------------------- //

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*time.Duration(shutdown_timeout),
	)
	defer cancel()

	if err := svc.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
