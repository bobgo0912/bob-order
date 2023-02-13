package repo

import (
	"github.com/bobgo0912/b0b-common/pkg/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

const OrderTableName = "order"

// Order order
type Order struct {
	Id         uint64    `db:"id" json:"id"`                  //p(PRI)
	Status     int       `db:"status" json:"status"`          //status  1.order 2.settle 3.cancel 4.error
	Period     string    `db:"period" json:"period"`          //period
	CardNumber int       `db:"card_number" json:"cardNumber"` //cardNumber
	CreateAt   time.Time `db:"create_at" json:"createAt"`     //create_at
	PlayerId   uint64    `db:"player_id" json:"playerId"`     //playerId
}

type OrderStore struct {
	*sql.BaseStore[Order]
}

func GetConnection() (*sqlx.DB, error) {
	if sql.OrderDb != nil {
		return sql.OrderDb, nil
	}
	var err error
	sql.OrderDb, err = sql.Db("order", nil)
	if err != nil {
		return nil, err
	}
	return sql.OrderDb, nil
}

func GetOrderStore() (*OrderStore, error) {
	connection, err := GetConnection()
	if err != nil {
		return nil, err
	}
	return &OrderStore{&sql.BaseStore[Order]{Db: connection, TableName: OrderTableName}}, nil
}
