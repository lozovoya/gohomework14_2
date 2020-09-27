package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Service struct {
	Pool *pgxpool.Pool
	Ctx  context.Context
}

type DbError struct {
	Err error
}

func NewDbError(err error) *DbError {
	return &DbError{Err: err}
}

func (e DbError) Error() string {
	return fmt.Sprintf("db error: %s", e.Err.Error())
}

func NewService(dsn string, ctx context.Context, pool *pgxpool.Pool) (*Service, error) {
	ctx = context.Background()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//defer pool.Close()
	return &Service{pool, ctx}, nil
}

func InitDb(ctx context.Context, dsn string) (pool *pgxpool.Pool) {

	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Println(err)
		return
	}
	//defer pool.Close()
	return pool
}
