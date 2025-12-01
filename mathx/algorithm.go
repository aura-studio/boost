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

func Count[T comparable](v T, s []T) int {
	return len(Positions(v, s))
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

// IntersectAny returns true if slice a and b share at least one common element.
// Time complexity: O(m+n), Space: O(min(m,n)).
func IntersectAny[T comparable](a, b []T) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	// Always build the map from the smaller slice.
	if len(a) > len(b) {
		a, b = b, a
	}
	m := make(map[T]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := m[v]; ok {
			return true
		}
	}
	return false
}

// IntersectCount returns the number of distinct common elements between a and b.
// Time complexity: O(m+n), Space: O(min(m,n)).
func IntersectCount[T comparable](a, b []T) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	if len(a) > len(b) {
		a, b = b, a
	}
	m := make(map[T]uint8, len(a))
	for _, v := range a {
		m[v] = 1 // mark presence in a
	}
	c := 0
	for _, v := range b {
		if m[v] == 1 { // first time seen in b
			c++
			m[v] = 2 // prevent double counting
		}
	}
	return c
}

// Intersect returns distinct intersection elements of a and b preserving
// the order they appear in b (first occurrence). If you need multiplicities,
// build a frequency map instead. Time: O(m+n), Space: O(min(m,n)).
func Intersect[T comparable](a, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	// Build from smaller slice for space efficiency.
	if len(a) > len(b) {
		a, b = b, a
	}
	seen := make(map[T]uint8, len(a))
	for _, v := range a {
		seen[v] = 1 // present in a
	}
	out := make([]T, 0)
	for _, v := range b {
		if seen[v] == 1 { // first time matched
			out = append(out, v)
			seen[v] = 2 // mark consumed to ensure distinctness
		}
	}
	return out
}
