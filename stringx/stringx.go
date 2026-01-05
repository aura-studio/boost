package stringx

import (
	"math"
	"strings"

	"github.com/aura-studio/boost/cast"
)

func Unique(ss []string) []string {
	m := make(map[string]struct{})
	for _, s := range ss {
		m[s] = struct{}{}
	}

	unique := make([]string, 0, len(m))
	for s := range m {
		unique = append(unique, s)
	}

	return unique
}

func Merge(ss ...string) string {
	var b strings.Builder
	b.Grow(64)
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

func PickLast(s string, sep string) string {
	lastIndex := strings.LastIndex(s, sep)
	if lastIndex < 0 {
		return s
	}
	return s[lastIndex+len(sep):]
}

func PruneLast(s string, sep string) string {
	lastIndex := strings.LastIndex(s, sep)
	if lastIndex < 0 {
		return s
	}
	return s[:lastIndex]
}

func PickFirst(s string, sep string) string {
	firstIndex := strings.Index(s, sep)
	if firstIndex < 0 {
		return s
	}
	return s[:firstIndex]
}

func PruneFirst(s string, sep string) string {
	firstIndex := strings.Index(s, sep)
	if firstIndex < 0 {
		return s
	}
	return s[firstIndex+len(sep):]
}

func ContainsAny(s string, v ...any) bool {
	size := len(v)
	if size == 0 {
		return false
	}

	if size > 1 {
		for _, item := range v {
			if strings.Contains(s, item.(string)) {
				return true
			}
		}
		return false
	}

	switch val := v[0].(type) {
	case string:
		return strings.Contains(s, val)
	case []string:
		for _, item := range val {
			if strings.Contains(s, item) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func Mod(s string, n int) int {
	var sum int
	for _, b := range []byte(s) {
		sum += int(b)
	}
	return sum % n
}

func CompareVersion(alpha string, beta string) int {
	if alpha == "" {
		alpha = cast.ToString(math.MinInt64)
	}
	if beta == "" {
		beta = cast.ToString(math.MinInt64)
	}
	alphaStrs := strings.Split(alpha, ".")
	betaStrs := strings.Split(beta, ".")

	var size int
	if len(alphaStrs) > len(betaStrs) {
		size = len(alphaStrs)
	} else {
		size = len(betaStrs)
	}

	alphaInts := make([]int, size)
	betaInts := make([]int, size)
	for i := 0; i < size; i++ {
		if i < len(alphaStrs) {
			alphaInts[i] = cast.ToInt(alphaStrs[i])
		} else {
			alphaInts[i] = math.MinInt64
		}
		if i < len(betaStrs) {
			betaInts[i] = cast.ToInt(betaStrs[i])
		} else {
			betaInts[i] = math.MinInt64
		}
	}

	for i := 0; i < size; i++ {
		if alphaInts[i] > betaInts[i] {
			return 1
		} else if alphaInts[i] < betaInts[i] {
			return -1
		}
	}

	return 0
}

func Shorten(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max]
}

func Capital(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
