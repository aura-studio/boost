package strick

import "strings"

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
