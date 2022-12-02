package device

import (
	"context"

	"github.com/aura-studio/boost/message"
)

type discard struct {
	*Base
}

func Discard() Device {
	return NewDiscard()
}

func NewDiscard() *discard {
	return &discard{
		Base: NewBase(),
	}
}

func (*discard) Process(context.Context, *message.Message) error {
	return nil
}
