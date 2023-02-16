package nats

import (
	"context"
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/b0b-common/pkg/nats"
	"github.com/bobgo0912/bob-settement/internal/constant"
	"github.com/bobgo0912/bob-settement/internal/model"
	"github.com/go-redis/redis/v9"
	nats2 "github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type Handler struct {
	Client *nats.JetClient
}

func NewHandler(client *nats.JetClient) *Handler {
	return &Handler{Client: client}
}

func (h *Handler) Start(ctx context.Context, client *redis.Client, queue chan *model.Prize) error {
	handle := NewHandle(client, h.Client.Client.Conn, queue)
	Register(constant.SettleBeginEvent, handle.GameBegin)
	Register(constant.SettleGoingEvent, handle.GameGoing)
	Register(constant.SettleFinishEvent, handle.GameFinish)
	Register(constant.SettleCancelEvent, handle.GameCancel)
	subscribe, err := h.Client.Client.Conn.Subscribe(constant.SettleEvent, handle.Handle)
	if err != nil {
		return errors.Wrap(err, "Subscribe fail")
	}
	go func(sub *nats2.Subscription) {
		for {
			select {
			case <-ctx.Done():
				sub.Unsubscribe()
				log.Info("nats Unsubscribe")
				return
			}
		}
	}(subscribe)
	return nil
}
