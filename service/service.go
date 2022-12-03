package service

import (
	"context"
	"log"
	"time"

	"github.com/aura-studio/boost/device"
	"github.com/aura-studio/boost/encoding"
	"github.com/aura-studio/boost/message"
	"github.com/aura-studio/boost/ref"
	"github.com/aura-studio/boost/route"
	"github.com/aura-studio/boost/safe"
	"github.com/aura-studio/boost/style"
)

type Service struct {
	Options
	client *device.Client
	router *device.Router
}

func New(a any, opts ...Option) *Service {
	client := device.NewClient("Client")
	device.Bus().Integrate(client)
	router := device.NewRouter(ref.TypeName(a)).Integrate(a)
	device.Bus().Integrate(router)

	s := &Service{
		client:  client,
		router:  router,
		Options: defaultOptions,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Service) Init() {}

func (s *Service) Invoke(routePath string, req string) (rsp string) {
	if err := safe.DoWithTimeout(10*time.Second, func(ctx context.Context) error {
		return s.client.Invoke(ctx, &message.Message{
			Route:    route.NewChainRoute(device.Addr(s.client), style.GoogleChain(routePath)),
			Encoding: encoding.NewJSON(),
			Data:     []byte(req),
		}, device.NewFuncProcessor(func(ctx context.Context, msg *message.Message) error {
			rsp = string(msg.Data)
			return nil
		}))
	}); err != nil {
		log.Panic(err)
	}
	return
}

func (s *Service) Close() {}
