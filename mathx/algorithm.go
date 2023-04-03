package mathx

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

func Positions[T comparable](v T, s []T) []int {
	pos := make([]int, 0)
	for p, n := range s {
		if n == v {
			pos = append(pos, p)
		}
	}
	return pos
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
