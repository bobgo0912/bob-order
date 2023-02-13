package nats

import (
	"context"
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/bobgo0912/b0b-common/pkg/nats"
	"github.com/bobgo0912/b0b-common/pkg/redis"
	"github.com/bobgo0912/bob-order/interal/constant"
	nats2 "github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"time"
)

type OrderHandler struct {
	Client *nats.JetClient
}

func NewOrderHandler(client *nats.JetClient) *OrderHandler {
	return &OrderHandler{Client: client}
}

func (o *OrderHandler) Start(ctx context.Context, redis *redis.Client) error {

	o.Client.Conn.AddStream(&nats2.StreamConfig{
		Name:        constant.OrderStream,
		Description: "order handler",
		Subjects:    []string{constant.OrderSubject},
		MaxAge:      time.Hour * 24 * 7,
	})
	handle := &Handle{RedisClient: redis}
	subscribe, err := o.Client.Conn.Subscribe(constant.OrderSubject, handle.Handle)
	if err != nil {
		log.Error("Subscribe fail err=", err)
		return errors.Wrap(err, "Subscribe fail")
	}
	go func(sub *nats2.Subscription) {
		for {
			select {
			case <-ctx.Done():
				sub.Unsubscribe()
				log.Info("nats Unsubscribe")
			}
		}
	}(subscribe)
	return nil
}
