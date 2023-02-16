package model

import "C"
import (
	"fmt"
)

type SettleCard struct {
	Id         uint64 `db:"id" json:"id"`                  //p(PRI)
	OrderId    uint64 `db:"order_id" json:"orderId"`       //orderId
	Numbers    string `db:"numbers" json:"numbers"`        //numbers
	Status     int    `db:"status" json:"status"`          //status 1.def 2.settle 3.cancel 4.error
	HandleNode string `db:"handle_node" json:"handleNode"` //
	Period     string `db:"period" json:"period"`          //period
	PlayerId   uint64 `db:"player_id" json:"playerId"`     //playerId
	NumbersAry []int
	Prize      []int
	Index      []int
}

func (c *SettleCard) ToPrize() *Prize {
	return &Prize{
		OrderId: c.OrderId,
		CardId:  c.Id,
		Prize:   fmt.Sprint(c.Prize),
		Period:  c.Period,
	}
}

func (c *SettleCard) GenIndex(i int) bool {
	i2 := i / 15
	if i%15 > 0 {
		i2 = i2 + 1
	}
	i2 = i2 - 1
	i3 := i2 * 5
	i4 := (i2+1)*5 - 1
	for j := i3; j <= i4; j++ {
		if i == c.NumbersAry[j] {
			c.Index = append(c.Index, j)
			return true
		}
	}
	return false
}
func (c *SettleCard) S(i int) bool {
	if c.Prize != nil {
		for _, i3 := range c.Prize {
			if i == i3 {
				return true
			}
		}
	}
	return false
}

func (c *SettleCard) Settle() {
	for i, pattern := range Patterns {
		if c.S(i) {
			continue
		}
		if c.F(pattern, c.Index) {
			c.Prize = append(c.Prize, i)
		}
	}
}
func (c *SettleCard) F(p, d []int) bool {
	var q int
	var l = len(p)
	for _, i := range p {
		if 12 == i {
			q++
			if q == l {
				return true
			}
			continue
		}
		for _, i2 := range d {
			if i == i2 {
				q++
				if q == l {
					return true
				}
				continue
			}
		}
	}
	return len(p) <= q
}
