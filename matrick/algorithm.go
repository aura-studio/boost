package matrick

import (
	"math/rand"
)

func FastFind[T1 Number](n T1, s []T1) int {
	var (
		min    = 0
		max    = len(s) - 1
		middle int
	)

	for {
		middle = (min + max) / 2
		if max == middle { // finish it: case min == max
			return middle
		} else if min == middle { // finish it: case min + 1 == max
			if n <= s[max] && s[max-1] < n {
				return max
			} else {
				return min
			}
		} else {
			if n <= s[middle-1] {
				max = middle - 1
			} else if s[middle] < n {
				min = middle + 1
			} else {
				return middle
			}
		}
	}
}

func Shuffle[T comparable](s []T) []T {
	length := int64(len(s))
	for i := length; i > 0; i-- {
		pos := rand.Int63n(i)
		s[i-1], s[pos] = s[pos], s[i-1]
	}
	return s
}

func Pos[T comparable](v T, s []T) []int {
	pos := make([]int, 0)
	for p, n := range s {
		if n == v {
			pos = append(pos, p)
		}
	}
	return pos
}

func Count[T comparable](v T, s []T) int {
	return len(Pos(v, s))
}

func In[T comparable](v T, s []T) bool {
	for _, n := range s {
		if v == n {
			return true
		}
	}

	return false
}

func Replace[T comparable](s []T, o T, n T) []T {
	for i := 0; i < len(s); i++ {
		if s[i] == o {
			s[i] = n
		}
	}
	return s
}

func Index[T comparable](s []T, v T) int {
	for i, n := range s {
		if n == v {
			return i
		}
	}
	return -1
}

func Sum[T Number](s []T) T {
	var sum T
	for _, n := range s {
		sum += n
	}
	return sum
}

// RandPR rands yes or no by probability
func RandPR(pr float64) bool {
	return rand.Float64() <= pr
}

// RandIntn Picks a random value in the specified object [s, e]
func RandIntn(s int, e int) int {
	return rand.Intn(e-s+1) + s
}

// RandIntn Picks a random value in the specified object [s, e]
func RandInt64n(s int64, e int64) int64 {
	return rand.Int63n(e-s+1) + s
}

func RandFloat64(s float64, e float64) float64 {
	r := s + rand.Float64()*(e-s)
	return r
}

func RandWeight[T Number](s []T) int {
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

// RandUnrepeated 以s为权重随机count个不重复的 返回索引的切片
func RandUnrepeated[T Number](s []T, count int) []int {
	result := make([]int, 0)
	for len(result) < count {
		var index int
		for index = RandWeight(s); In(index, result); {
			index = RandWeight(s)
		}
		result = append(result, index)
	}
	return result
}

// MaxContinuousCount returns max continuous count in slice  [start end) index
func MaxContinuousCount[T comparable](v T, s []T) (int, int, int) {
	maxCount := 0
	start := 0
	end := 0
	for index := 0; index < len(s); {
		if s[index] == v {
			tmp := index + 1
			for tmp < len(s) && s[tmp] == v {
				tmp++
			}
			if tmp-index > maxCount {
				maxCount = tmp - index
				start = index
				end = tmp
			}
			index = tmp
		} else {
			index++
		}
	}
	return maxCount, start, end
}

func ContinuousPositions[T comparable](v T, s []T) [][]int {
	rt := make([][]int, 0)
	for index := 0; index < len(s); {
		if s[index] == v {
			rtt := make([]int, 0)
			rtt = append(rtt, index)
			tmp := index + 1
			for tmp < len(s) && s[tmp] == v {
				rtt = append(rtt, tmp)
				tmp++
			}
			index = tmp
			rt = append(rt, rtt)
		} else {
			index++
		}
	}
	return rt
}
