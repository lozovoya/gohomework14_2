package card

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Card struct {
	Id      int64
	Number  string
	Balance int64
	Issuer  string
	Status  string
}

type Transaction struct {
	Id          int64
	Amount      int64
	Category    int64
	Description int64
	Logo        int64
}

type Service struct {
	Cards       []*Card
	Transaction []*Transaction
}

func NewService() *Service { return &Service{} }

func GetCards(ctx context.Context, pool *pgxpool.Pool, id int) (cards []*Card) {

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(ctx,
		"SELECT id, number, balance, issuer, status FROM cards WHERE owner_id = $1 LIMIT 50", id)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		log.Println(rows.Values())
		card := &Card{}
		err = rows.Scan(&card.Id,
			&card.Number,
			&card.Balance,
			&card.Issuer,
			&card.Status)
		cards = append(cards, card)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return
	}
	return cards
}

func GetTransactions(ctx context.Context, pool *pgxpool.Pool, id int) (transactions []*Transaction) {

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(
		ctx,
		"SELECT t.id, t.amount, t.category_id, t.description_id, t.logo_id FROM transactions t JOIN categories ON t.category_id = categories.id JOIN descriptions ON  t.description_id = descriptions.id JOIN logos ON t.logo_id = logos.id    WHERE t.card_id = $1 LIMIT 50",
		id,
	)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		log.Println(rows.Values())
		transaction := &Transaction{}
		err = rows.Scan(&transaction.Id,
			&transaction.Amount,
			&transaction.Category,
			&transaction.Description,
			&transaction.Logo)
		transactions = append(transactions, transaction)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return
	}
	return transactions
}
