package mathx

import (
	"math/rand"

	"github.com/aura-studio/boost/mathx"
)

type Signed interface {
	int | int8 | int16 | int32 | int64
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

type Rand struct {
	*rand.Rand
}

func NewRand(seed int64) *Rand {
	return &Rand{
		Rand: rand.New(rand.NewSource(seed)),
	}
}

// PR rands yes or no by probability
func PR(rand *Rand, pr float64) bool {
	return rand.Float64() <= pr
}

func (rand *Rand) PR(pr float64) bool {
	return rand.Float64() <= pr
}

// Intn picks a random value in the specified ramge [0, n)
func Intn(rand *Rand, n int) int {
	return rand.Intn(n)
}

// Intn picks a random value in the specified ramge [0, n)
func (rand *Rand) Intn(n int) int {
	return rand.Rand.Intn(n)
}

// RangeIntn picks a random value in the specified ramge [s, e]
func RangeIntn(rand *Rand, s int, e int) int {
	return rand.Intn(e-s+1) + s
}

// RangeIntn picks a random value in the specified ramge [s, e]
func (rand *Rand) RangeIntn(s int, e int) int {
	return rand.Intn(e-s+1) + s
}

// Int63n picks a random value in the specified ramge [0, n)
func Int63n(rand *Rand, n int64) int64 {
	return rand.Int63n(n)
}

// Int63n picks a random value in the specified ramge [0, n)
func (rand *Rand) Int63n(n int64) int64 {
	return rand.Rand.Int63n(n)
}

// RangeInt63n picks a random value in the specified ramge [s, e]
func RangeInt63n(rand *Rand, s int64, e int64) int64 {
	return rand.Int63n(e-s+1) + s
}

// RangeInt63n picks a random value in the specified ramge [s, e]
func (rand *Rand) RangeInt63n(s int64, e int64) int64 {
	return rand.Int63n(e-s+1) + s
}

// Float64 picks a random value in the specified ramge [0, 1)
func Float64(rand *Rand) float64 {
	return rand.Float64()
}

// Float64 picks a random value in the specified ramge [0, 1)
func (rand *Rand) Float64() float64 {
	return rand.Rand.Float64()
}

// RangeFloat64 picks a random value in the specified range [s, e)
func RangeFloat64(rand *Rand, s float64, e float64) float64 {
	r := s + rand.Float64()*(e-s)
	return r
}

// RangeFloat64 picks a random value in the specified range [s, e)
func (rand *Rand) RangeFloat64(s float64, e float64) float64 {
	r := s + rand.Float64()*(e-s)
	return r
}

// RandWeight picks a random value in the specified slice by weight
func RandWeight[T Number](rand *Rand, s []T) int {
	var weightSum float64
	for _, r := range s {
		weightSum += float64(r)
	}
	n := rand.Float64() * weightSum
	var lastKey int
	for i, r := range s {
		nReach := float64(r)
		if n <= nReach {
			return i
		}
		n -= nReach
		lastKey = i
	}
	return lastKey
}

// RandWeightMap picks a random value in the specified map by weight
func RandWeightMap[T1 comparable, T2 Number](rand *Rand, m map[T1]T2) T1 {
	var keys = make([]T1, 0, len(m))
	var values = make([]T2, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	i := RandWeight(rand, values)
	return keys[i]
}

// RandUnrepeated picks unrepeated random values in the specified object by weight
func RandUnrepeated[T Number](rand *Rand, s []T, count int) []int {
	result := make([]int, 0)
	for len(result) < count {
		var index int
		for index = RandWeight(rand, s); mathx.In(index, result); {
			index = RandWeight(rand, s)
		}
		result = append(result, index)
	}
	return result
}

// RandShuffle shuffles the specified slice
func RandShuffle[T comparable](rand *Rand, s []T) []T {
	length := int64(len(s))
	for i := length; i > 0; i-- {
		pos := rand.Int63n(i)
		s[i-1], s[pos] = s[pos], s[i-1]
	}
	return s
}
