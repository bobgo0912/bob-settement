package model

import (
	"github.com/bobgo0912/b0b-common/pkg/sql"
	"github.com/jmoiron/sqlx"
)

const PrizeTableName = "prize"

// Prize prize
type Prize struct {
	Id      uint64 `db:"id" json:"id"`            //id(PRI)
	OrderId uint64 `db:"order_id" json:"orderId"` //orderId
	CardId  uint64 `db:"card_id" json:"cardId"`   //cardId
	Prize   string `db:"prize" json:"prize"`      //prize
	Period  string `db:"period" json:"period"`    //period
}

type PrizeStore struct {
	*sql.BaseStore[Prize]
}

func GetConnection() (*sqlx.DB, error) {
	if SettleDb != nil {
		return SettleDb, nil
	}
	var err error
	SettleDb, err = sql.Db("settle", nil)
	if err != nil {
		return nil, err
	}
	return SettleDb, nil
}

func GetPrizeStore() (*PrizeStore, error) {
	connection, err := GetConnection()
	if err != nil {
		return nil, err
	}
	return &PrizeStore{&sql.BaseStore[Prize]{Db: connection, TableName: PrizeTableName}}, nil
}
