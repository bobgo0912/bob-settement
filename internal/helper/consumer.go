package helper

import (
	"context"
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/bob-settement/internal/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type Type []*model.Prize

type MysqlConsumer struct {
	EventQueue     chan *model.Prize
	BatchSize      int
	Workers        int
	LingerTime     time.Duration
	BatchProcessor func(batch Type) error
	ErrHandler     func(err error, batch Type)
}

func NewMysqlConsumer(db *model.PrizeStore) *MysqlConsumer {
	return &MysqlConsumer{
		EventQueue: make(chan *model.Prize, 10000),
		BatchSize:  500,
		Workers:    2,
		LingerTime: 50 * time.Millisecond,
		BatchProcessor: func(batch Type) error {
			if len(batch) < 0 {
				return nil
			}
			defer func() {
				if r := recover(); r != nil {
					log.Error("recover err=", r)
				}
			}()
			err := db.MultipleInsert(context.TODO(), batch)
			if err != nil {
				log.Error("MultipleInsert fail err=", err)
				return errors.Wrap(err, "MultipleInsert fail")
			}
			return nil
		},
		ErrHandler: func(err error, batch Type) {
			log.Error("consumer type=", batch, " err=", err.Error())
		},
	}
}

func (c *MysqlConsumer) Start(ctx context.Context) {
	for i := 0; i < c.Workers; i++ {
		go func() {
			var batch Type
			lingerTimer := time.NewTimer(0)
			if !lingerTimer.Stop() {
				<-lingerTimer.C
			}
			defer lingerTimer.Stop()

			for {
				select {
				case msg := <-c.EventQueue:
					batch = append(batch, msg)
					if len(batch) != c.BatchSize {
						if len(batch) == 1 {
							lingerTimer.Reset(c.LingerTime)
						}
						break
					}
					if err := c.BatchProcessor(batch); err != nil {
						c.ErrHandler(err, batch)
					}
					if !lingerTimer.Stop() {
						<-lingerTimer.C
					}
					batch = make(Type, 0)
				case <-lingerTimer.C:
					if err := c.BatchProcessor(batch); err != nil {
						c.ErrHandler(err, batch)
					}
					batch = make(Type, 0)
				case <-ctx.Done():
					zap.S().Info("MysqlConsumer done")
					return
				}
			}
		}()
	}

}
