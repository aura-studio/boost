package style

import (
	"strings"

	"github.com/aura-studio/boost/magic"
)

func Standardize(s string, sep magic.SeparatorType) string {
	if s == "" || sep == magic.SeparatorLazy {
		return s
	}

	// sanitize: keep only letters, digits and separator
	s = sanitize(s, sep)

	var words []string
	if sep == magic.SeparatorNone {
		words = []string{s}
	} else {
		words = strings.Split(s, sep)
	}

	var sb strings.Builder
	sb.Grow(len(s))
	for _, word := range words {
		if word == "" {
			continue
		}
		r := []rune(word)
		for i := range r {
			r[i] = toLower(r[i])
		}
		if len(r) > 0 && !isUpper(r[0]) {
			if r[0] >= 'a' && r[0] <= 'z' {
				r[0] = toUpper(r[0])
			}
		}
		sb.WriteString(string(r))
	}
	return sb.String()
}

func Unstandardize(s string, sep magic.SeparatorType) string {
	if s == "" || sep == magic.SeparatorLazy {
		return s
	}

	// sanitize: keep only letters, digits and separator
	s = sanitize(s, sep)
	// Split into words considering consecutive uppercase sequences (acronyms)
	// Inline splitStandardWords logic
	var words []string
	r := []rune(s)
	if len(r) > 0 {
		start := 0
		i := 0
		for i < len(r) {
			if isUpper(r[i]) {
				j := i + 1
				for j < len(r) && isUpper(r[j]) {
					j++
				}
				if j < len(r) && !isUpper(r[j]) {
					if j-i > 1 {
						words = append(words, string(r[i:j-1]))
						i = j - 1
						k := i + 1
						for k < len(r) && !isUpper(r[k]) {
							k++
						}
						words = append(words, string(r[i:k]))
						i = k
						start = i
						continue
					}
					k := j + 1
					for k < len(r) && !isUpper(r[k]) {
						k++
					}
					words = append(words, string(r[i:k]))
					i = k
					start = i
					continue
				}
				words = append(words, string(r[i:j]))
				i = j
				start = i
				continue
			} else {
				j := i + 1
				for j < len(r) && !isUpper(r[j]) {
					j++
				}
				words = append(words, string(r[i:j]))
				i = j
				start = i
				continue
			}
		}
		if len(words) == 0 && start < len(r) {
			words = append(words, string(r[start:]))
		}
	}

	// normalize to lowercase words
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return strings.Join(words, sep)
}
