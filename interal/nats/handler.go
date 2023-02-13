package nats

import (
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/b0b-common/pkg/redis"
	"github.com/nats-io/nats.go"
	"sync"
)

type HandleFunc func(data []byte) error

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
}

func (s *Handle) Handle(msg *nats.Msg) {
	value, ok := GetHandlers(msg.Subject)
	if ok {
		if len(msg.Data) < 1 {
			log.Error("bad msg=", msg)
			return
		}
		err := value(msg.Data)
		if err != nil {
			log.Error("Handle err=", err.Error())
		}
	} else {
		log.Warn("Handle GetHandlers fail Sub=", msg.Subject)
	}
}
