package radix

import (
	"math"
)

type Radix struct {
	*Charset
	base    int
	maxSize int // -1 means no limit
}

func New(maxSize int, c *Charset) *Radix {
	r := &Radix{
		Charset: c,
		base:    c.Size(),
		maxSize: maxSize,
	}
	return r
}

func (r *Radix) Encode(n uint64) []byte {
	n = r.validate(n)
	data := make([]byte, 0)
	for n > 0 {
		data = append(data, r.charset[n%uint64(r.base)])
		n /= uint64(r.base)
	}
	return data
}

func (r *Radix) Decode(data []byte) uint64 {
	var n uint64
	for i := len(data) - 1; i >= 0; i-- {
		n = n*uint64(r.base) + uint64(r.charsetMap[data[i]])
	}
	return r.validate(n)
}

func (r *Radix) validate(i uint64) uint64 {
	if r.maxSize > 0 {
		return i % uint64(math.Pow(float64(r.base), float64(r.maxSize)))
	}
	return i
}
