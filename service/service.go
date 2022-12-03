package service

import (
	"context"
	"log"
	"reflect"
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
	target any
	client *device.Client
	router *device.Router
	init   func()
	close  func()
}

func New(target any, opts ...Option) *Service {
	t := reflect.TypeOf(target)
	if !(t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct) {
		log.Panic("service must be a pointer to struct")
	}

	client := device.NewClient("Client")
	device.Bus().Integrate(client)
	router := device.NewRouter(ref.TypeName(target)).Integrate(target)
	device.Bus().Integrate(router)

	s := &Service{
		Options: defaultOptions,
		target:  target,
		client:  client,
		router:  router,
	}

	if init, ok := t.MethodByName("Init"); ok {
		if init.Type.NumIn() == 1 && init.Type.NumOut() == 0 {
			s.init = func() {
				init.Func.Call([]reflect.Value{reflect.ValueOf(target)})
			}
		}
	}

	if close, ok := t.MethodByName("Close"); ok {
		if close.Type.NumIn() == 1 && close.Type.NumOut() == 0 {
			s.close = func() {
				close.Func.Call([]reflect.Value{reflect.ValueOf(target)})
			}
		}
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Service) Init() {
	s.init()
}

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

func (s *Service) Close() {
	s.close()
}
