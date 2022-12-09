package radix

import (
	"math"
	"math/rand"
)

type Radix struct {
	base      int
	maxSize   int
	seed      int64
	charset   []byte
	reCharset map[byte]int
}

func New(maxSize int, seed int64, charset []byte) *Radix {
	r := &Radix{
		base:      len(charset),
		maxSize:   maxSize,
		seed:      seed,
		charset:   charset,
		reCharset: make(map[byte]int),
	}
	r.init()
	return r
}

func (r *Radix) init() {
	rand.Seed(r.seed)
	rand.Shuffle(len(r.charset), func(i, j int) {
		r.charset[i], r.charset[j] = r.charset[j], r.charset[i]
	})
	for i, b := range r.charset {
		r.reCharset[b] = i
	}
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
		n = n*uint64(r.base) + uint64(r.reCharset[data[i]])
	}
	return r.validate(n)
}

func (r *Radix) validate(i uint64) uint64 {
	return i % uint64(math.Pow(float64(r.base), float64(r.maxSize)))
}

// var Radix64 *Radix

// func init() {
// 	t, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
// 	base64 = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_=")
// 	Raidx64 = New(64, 8, []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_="), t.Unix())
// }

// var reverseBase64 = map[byte]uint8{}

// func init() {
// 	rand.Seed(t.Unix())
// 	rand.Shuffle(len(base64), func(i, j int) {
// 		base64[i], base64[j] = base64[j], base64[i]
// 	})
// 	for i, b := range base64 {
// 		reverseBase64[b] = uint8(i)
// 	}
// }

// func base64Encode(i uint64) []byte {
// 	result := make([]byte, 0)
// 	for i > 0 {
// 		result = append(result, base64[i%64])
// 		i /= 64
// 	}
// 	return result
// }

// func base64Decode(b []byte) uint64 {
// 	var result uint64
// 	for i := len(b) - 1; i >= 0; i-- {
// 		result = result*64 + uint64(reverseBase64[b[i]])
// 	}
// 	return result
// }

// func base64Validate(i uint64, maxLen uint8) uint64 {
// 	return i % uint64(math.Pow(64, float64(maxLen)))
// }
