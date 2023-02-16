package model

import (
	"github.com/bobgo0912/b0b-common/pkg/config"
	"strconv"
	"strings"
)

type Card struct {
	Id         uint64 `db:"id" json:"id"`                  //p(PRI)
	OrderId    uint64 `db:"order_id" json:"orderId"`       //orderId
	Numbers    string `db:"numbers" json:"numbers"`        //numbers
	Status     int    `db:"status" json:"status"`          //status 1.def 2.settle 3.cancel 4.error
	HandleNode string `db:"handle_node" json:"handleNode"` //
	Period     string `db:"period" json:"period"`          //period
	PlayerId   uint64 `db:"player_id" json:"playerId"`     //playerId
}

func (c *Card) ToSettle() *SettleCard {
	return &SettleCard{
		Id:         c.Id,
		OrderId:    c.OrderId,
		Numbers:    c.Numbers,
		Status:     c.Status,
		HandleNode: config.Cfg.NodeId,
		Period:     c.Period,
		PlayerId:   c.PlayerId,
		NumbersAry: c.Str2Numbers(),
	}
}

func (c *Card) Str2Numbers() []int {
	split := strings.Split(c.Numbers, ",")
	if len(split) < 1 {
		return nil
	}
	ints := make([]int, 0, 25)
	for i, s := range split {
		if i == 12 {
			ints = append(ints, 0)
			continue
		}
		atoi, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		ints = append(ints, atoi)
	}
	return ints
}
