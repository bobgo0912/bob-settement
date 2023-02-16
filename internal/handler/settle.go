package handler

import (
	"context"
	"github.com/bobgo0912/b0b-common/pkg/log"
	gm "github.com/bobgo0912/bob-armory/pkg/game/model"
	"github.com/bobgo0912/bob-settement/internal/model"
	sm "github.com/bobgo0912/bob-settement/internal/model"
	"github.com/bobgo0912/bob-settement/internal/util"
	"github.com/pkg/errors"
)

type Settle struct {
	StartPrize  bool
	Period      string
	Cards       []*model.SettleCard
	PrizeNumber []int
	PrizeQueue  chan *sm.Prize
}

func (s *Settle) Add(card *model.SettleCard) error {
	if s.Period != "" {
		return errors.New("period is blank")
	}
	s.Cards = append(s.Cards, card)
	return nil
}

func (s *Settle) Prize(m *gm.GameRoundEvent) error {
	number, err := util.GetLastPrizeNumber(m.PrizeNumber)
	if err != nil {
		return errors.Wrap(err, "GetLastPrizeNumber fail")
	}
	s.PrizeNumber = append(s.PrizeNumber, number)
	return s.ExtraPatternsSettle(number)
}

func (s *Settle) Finish(m *gm.GameRoundEvent) error {
	//todo verify
	ids := make([]uint64, 0, len(s.Cards))
	cards := make([]*model.SettleCard, 0)
	for _, card := range s.Cards {
		ids = append(ids, card.Id)
		if len(card.Prize) > 0 {
			cards = append(cards, card)
		}
	}

	store, err := model.GetPrizeStore()
	if err != nil {
		log.Error("GetPrizeStore fail err=", err)
		return err
	}
	prizes := make([]*model.Prize, 0, len(cards))
	for _, card := range cards {
		prizes = append(prizes, card.ToPrize())
	}

	err = store.MultipleInsert(context.TODO(), prizes)
	if err != nil {
		log.Error("MultipleInsert fail err=", err)
		return errors.Wrap(err, "MultipleInsert fail")
	}

	//s.PrizeQueue <- cards
	s.Period = ""
	s.Cards = make([]*model.SettleCard, 0)
	s.PrizeNumber = nil

	return nil
}
