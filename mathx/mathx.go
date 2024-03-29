package mathx

import "time"

type Signed interface {
	int | int8 | int16 | int32 | int64 | time.Duration
}

type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}
