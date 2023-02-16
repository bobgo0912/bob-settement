package handler

import (
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/bob-settement/internal/helper"
	"github.com/bobgo0912/bob-settement/internal/model"
	"github.com/pkg/errors"
)

func (s *Settle) ExtraPatternsSettle(number int) error {
	if len(s.Cards) < 1 {
		return nil
	}
	array, err := helper.SplitArray[*model.SettleCard](s.Cards, 10000)
	if err != nil {
		log.Error("SplitArray fail err=", err)
		return errors.Wrap(err, "SplitArray fail")
	}
	for _, ts := range array {
		go func(a []*model.SettleCard, num int) {
			for _, card := range a {
				if card.GenIndex(num) {
					card.Settle()
				}
			}
		}(ts, number)
	}
	return nil
}
