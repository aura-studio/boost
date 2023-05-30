package stringx

import (
	"math"
	"strings"

	"github.com/aura-studio/boost/cast"
)

// Merge joins all strings
func Merge(ss ...string) string {
	var b strings.Builder
	b.Grow(64)
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

// PickLast splits s with sep, and get last piece
func PickLast(s string, sep string) string {
	lastIndex := strings.LastIndex(s, sep)
	if lastIndex < 0 {
		return s
	}
	return s[lastIndex+len(sep):]
}

// PruneLast splits s with sep, and get all pieces except the last
func PruneLast(s string, sep string) string {
	lastIndex := strings.LastIndex(s, sep)
	if lastIndex < 0 {
		return s
	}
	return s[:lastIndex]
}

// PickFirst splits s with sep, and get first piece
func PickFirst(s string, sep string) string {
	firstIndex := strings.Index(s, sep)
	if firstIndex < 0 {
		return s
	}
	return s[:firstIndex]
}

// PruneFirst splits s with sep, and get all pieces except the first
func PruneFirst(s string, sep string) string {
	firstIndex := strings.Index(s, sep)
	if firstIndex < 0 {
		return s
	}
	return s[firstIndex+len(sep):]
}

// Mod returns the remainder of s divided by n
func Mod(s string, n int) int {
	var sum int
	for _, b := range []byte(s) {
		sum += int(b)
	}
	return sum % n
}

// CompareVersion compares two version strings
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
