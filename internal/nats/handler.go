package nats

import (
	"encoding/json"
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/bob-armory/pkg/game/model"
	"github.com/bobgo0912/bob-settement/internal/handler"
	sm "github.com/bobgo0912/bob-settement/internal/modle"
	"github.com/go-redis/redis/v9"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"sync"
)

type HandleFunc func(msg *nats.Msg) error

var (
	handlers        = make(map[string]HandleFunc)
	handlersRWMutex sync.RWMutex
)

func Register(key string, value HandleFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()
	handlers[key] = value

	return
}

func GetHandlers(key string) (value HandleFunc, ok bool) {
	handlersRWMutex.RLock()
	defer handlersRWMutex.RUnlock()
	value, ok = handlers[key]
	return
}

type Handle struct {
	RedisClient *redis.Client
	JetClient   *nats.Conn
	Handler     *handler.Settle
}

func NewHandle(rc *redis.Client, jc *nats.Conn) *Handle {
	return &Handle{RedisClient: rc, JetClient: jc, Handler: &handler.Settle{Cards: make([]*sm.SettleCard, 0)}}
}

func (s *Handle) Handle(msg *nats.Msg) {
	value, ok := GetHandlers(msg.Subject)
	if ok {
		if len(msg.Data) < 1 {
			log.Error("bad msg=", msg)
			return
		}
		err := value(msg)
		if err != nil {
			log.Error("Handle err=", err.Error())
		}
	} else {
		log.Warn("Handle GetHandlers fail Sub=", msg.Subject)
	}
}

func (s *Handle) GameBegin(msg *nats.Msg) error {
	data := msg.Data
	if s.Handler.Period != "" {
		//todo handle error
	}
	var event model.GameRoundEvent
	err := json.Unmarshal(data, &event)
	if err != nil {
		return errors.Wrap(err, "Unmarshal")
	}
	if event.GameRoundId == "" {
		return errors.New("bad period")
	}
	s.Handler.Period = event.GameRoundId
	return nil
}

func (s *Handle) AddCard(msg *nats.Msg) error {
	data := msg.Data
	var c sm.Card
	err := json.Unmarshal(data, &c)
	if err != nil {
		return errors.Wrap(err, "Unmarshal fail")
	}
	err = s.Handler.Add(c.ToSettle())
	if err != nil {
		msg.Respond([]byte("0"))
		return errors.Wrap(err, "add fail")
	}
	msg.Respond([]byte("1"))
	return err
}

func (s *Handle) GameGoing(msg *nats.Msg) error {
	data := msg.Data
	var gameEvent model.GameRoundEvent
	err := json.Unmarshal(data, &gameEvent)
	if err != nil {
		log.Error("Unmarshal err=", err.Error())
		return errors.Wrap(err, "Unmarshal fail")
	}
	if s.Handler.Period == "" {
		s.Handler.Period = gameEvent.GameRoundId
	} else {
		if s.Handler.Period != gameEvent.GameRoundId {
			return errors.New("period not same")
		}
	}
	return s.Handler.Prize(&gameEvent)
}

func (s *Handle) GameFinish(msg *nats.Msg) error {
	data := msg.Data
	var gameEvent model.GameRoundEvent
	err := json.Unmarshal(data, &gameEvent)
	if err != nil {
		log.Error("Unmarshal err=", err.Error())
		return errors.Wrap(err, "Unmarshal fail")
	}
	return s.Handler.Finish(&gameEvent)
}

func (s *Handle) GameCancel(msg *nats.Msg) error {
	return nil
}
