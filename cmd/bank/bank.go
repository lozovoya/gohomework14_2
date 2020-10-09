package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lozovoya/gohomework14_2/cmd/bank/app"
	"github.com/lozovoya/gohomework14_2/pkg/card"
	"log"
	"net"
	"net/http"
	"os"
)

const defaultPort = "9999"
const defaultHost = "0.0.0.0"
const dbcon = "postgres://app:pass@bankdb:5432/db"

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	log.Println(host)
	log.Println(port)

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(addr string) (err error) {

	mux := http.NewServeMux()
	dsn := dbcon
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println(err)
		return err
	}
	defer pool.Close()
	cardSvc := card.NewService(pool)

	application := app.NewServer(mux, cardSvc)

	application.Init()
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return server.ListenAndServe()

}
