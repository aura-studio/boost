package style

import (
	"strings"

	"github.com/aura-studio/boost/magic"
)

var (
	googleChain = *NewChainStyle(magic.SeparatorSlash, magic.SeparatorHyphen)
	unixChain   = *NewChainStyle(magic.SeparatorPeriod, magic.SeparatorUnderscore)
)

func camelize(s string) string {
	s = strings.ToLower(s)
	b := []byte(s)
	if b[0] >= 'a' && b[0] <= 'z' {
		b[0] -= 32
	}
	return string(b)
}

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

	var b = []byte{}
	for _, word := range words {
		word = camelize(word)
		b = append(b, []byte(word)...)
	}
	return string(b)
}

// Unstandardize does the inverse of Standardize:
// it splits a Standardized string (PascalCase with known abbreviations)
// into lowercase words joined by the provided separator.
func Unstandardize(s string, sep magic.SeparatorType) string {
	if s == "" || sep == magic.SeparatorLazy {
		return s
	}

	// sanitize: keep only letters, digits and separator
	s = sanitize(s, sep)

	// Split into words considering consecutive uppercase sequences (acronyms)
	words := splitStandardWords(s)

	// normalize to lowercase words
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return strings.Join(words, sep)
}

// splitStandardWords splits a Standardized string into words,
// preserving acronym blocks like HTTP, JSON, UUID.
func splitStandardWords(s string) []string {
	if s == "" {
		return []string{}
	}
	// work with runes for safety
	r := []rune(s)
	var words []string

	start := 0
	i := 0
	for i < len(r) {
		// advance through current segment
		if isUpper(r[i]) {
			// consume consecutive uppers
			j := i + 1
			for j < len(r) && isUpper(r[j]) {
				j++
			}
			if j < len(r) && isLower(r[j]) {
				// Case like JSONData: keep JSON as one word, then Data
				// However, we may have j-i >= 2; to avoid splitting D from Data,
				// we cut before the last upper so that Data remains capitalized
				// Example: "JSONData" -> [JSON] [Data]
				if j-i > 1 {
					// word = r[i:j-1], then continue from j-1
					words = append(words, string(r[i:j-1]))
					i = j - 1
					// continue to process mixed-case word starting at i
					// consume until next upper
					k := i + 1
					for k < len(r) && isLower(r[k]) {
						k++
					}
					words = append(words, string(r[i:k]))
					i = k
					start = i
					continue
				}
				// Single upper followed by lower: start of regular word
				k := j + 1
				for k < len(r) && isLower(r[k]) {
					k++
				}
				words = append(words, string(r[i:k]))
				i = k
				start = i
				continue
			}
			// All uppers till end or followed by upper: acronym block
			words = append(words, string(r[i:j]))
			i = j
			start = i
			continue
		} else {
			// start with lowercase: consume until next upper
			j := i + 1
			for j < len(r) && isLower(r[j]) {
				j++
			}
			words = append(words, string(r[i:j]))
			i = j
			start = i
			continue
		}
	}
	// fallback if nothing parsed
	if len(words) == 0 && start < len(r) {
		words = append(words, string(r[start:]))
	}
	return words
}

func isUpper(r rune) bool { return r >= 'A' && r <= 'Z' }

func isLower(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

// sanitize removes all characters except letters, digits and provided separator(s).
func sanitize(s string, sep magic.SeparatorType) string {
	if s == "" {
		return s
	}
	// build whitelist for separators (can be more than one char)
	sepRunes := map[rune]struct{}{}
	for _, sr := range []rune(sep) {
		sepRunes[sr] = struct{}{}
	}
	r := []rune(s)
	out := make([]rune, 0, len(r))
	for _, ch := range r {
		if isUpper(ch) || isLower(ch) {
			out = append(out, ch)
			continue
		}
		if _, ok := sepRunes[ch]; ok {
			out = append(out, ch)
		}
		// else: drop
	}
	return string(out)
}

type ChainStyle struct {
	ChainSeperator magic.SeparatorType
	WordSeparator  magic.SeparatorType
}

func NewChainStyle(chainSeparator, wordSeparator string) *ChainStyle {
	return &ChainStyle{
		ChainSeperator: chainSeparator,
		WordSeparator:  wordSeparator,
	}
}

func Chain(s string, cs ChainStyle) []string {
	return cs.Chain(s)
}

func (cs ChainStyle) Chain(s string) []string {
	chain := strings.Split(s, cs.ChainSeperator)
	for index := 0; index < len(chain); index++ {
		chain[index] = Standardize(chain[index], cs.WordSeparator)
	}
	return chain
}

func GoogleChain(s string) []string {
	return Chain(s, googleChain)
}

func UnixChain(s string) []string {
	return Chain(s, unixChain)
}
