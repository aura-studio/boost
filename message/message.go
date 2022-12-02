package message

import (
	"github.com/aura-studio/boost/encoding"
	"github.com/aura-studio/boost/route"
)

type Message struct {
	ID       uint64
	Route    route.Route
	Encoding encoding.Encoding
	Data     []byte
}
