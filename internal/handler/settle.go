package handler

import (
	"github.com/bobgo0912/bob-armory/pkg/game/model"
	"github.com/bobgo0912/bob-settement/internal/modle"
	"github.com/bobgo0912/bob-settement/internal/util"
	"github.com/pkg/errors"
)

type Settle struct {
	StartPrize  bool
	Period      string
	Cards       []*modle.SettleCard
	PrizeNumber []int
}

func (s *Settle) Add(card *modle.SettleCard) error {
	if s.Period != "" {
		return errors.New("period is blank")
	}
	s.Cards = append(s.Cards, card)
	return nil
}

func (s *Settle) Prize(m *model.GameRoundEvent) error {
	number, err := util.GetLastPrizeNumber(m.PrizeNumber)
	if err != nil {
		return errors.Wrap(err, "GetLastPrizeNumber fail")
	}
	s.PrizeNumber = append(s.PrizeNumber, number)
	return s.ExtraPatternsSettle(number)
}

func (s *Settle) Finish(m *model.GameRoundEvent) error {
	//todo verify
	ids := make([]uint64, 0, len(s.Cards))
	cards := make([]*modle.SettleCard, 0)

	for _, card := range s.Cards {
		ids = append(ids, card.Id)
		if len(card.Prize) > 0 {
			cards = append(cards, card)
		}
	}

	return nil
}
