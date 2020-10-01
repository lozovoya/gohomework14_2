package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Service struct {
	Pool *pgxpool.Pool
	Ctx  context.Context
}

func NewService(dsn string) (*Service, error) {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//defer pool.Close()
	return &Service{pool, ctx}, nil
}
