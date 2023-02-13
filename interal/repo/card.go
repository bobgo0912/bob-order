package repo

import (
	"github.com/bobgo0912/b0b-common/pkg/sql"
)

const CardTableName = "card"

// Card card
type Card struct {
	Id         uint64 `db:"id" json:"id"`                  //p(PRI)
	OrderId    uint64 `db:"order_id" json:"orderId"`       //orderId
	Numbers    string `db:"numbers" json:"numbers"`        //numbers
	Status     int    `db:"status" json:"status"`          //status 1.def 2.settle 3.cancel 4.error
	HandleNode string `db:"handle_node" json:"handleNode"` //
	Period     string `db:"period" json:"period"`          //period
	PlayerId   uint64 `db:"player_id" json:"playerId"`     //playerId

}
type CardStore struct {
	*sql.BaseStore[Card]
}

func GetCardStore() (*CardStore, error) {
	connection, err := GetConnection()
	if err != nil {
		return nil, err
	}
	return &CardStore{&sql.BaseStore[Card]{Db: connection, TableName: CardTableName}}, nil
}
