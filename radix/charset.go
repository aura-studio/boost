package radix

import (
	"math/rand"
)

type Charset struct {
	seed          int64
	originCharset []byte
	charset       []byte
	charsetMap    map[byte]int
}

func NewCharset(seed int64, charset []byte) *Charset {
	c := &Charset{
		seed:          seed,
		originCharset: make([]byte, len(charset)),
		charset:       make([]byte, len(charset)),
		charsetMap:    make(map[byte]int),
	}
	copy(c.originCharset, charset)
	copy(c.charset, charset)

	r := rand.New(rand.NewSource(c.seed))
	r.Shuffle(len(c.charset), func(i, j int) {
		c.charset[i], c.charset[j] = c.charset[j], c.charset[i]
	})

	for i, b := range c.charset {
		c.charsetMap[b] = i
	}

	return c
}

func (c *Charset) Size() int {
	return len(c.charset)
}
