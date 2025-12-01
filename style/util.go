package style

import "github.com/aura-studio/boost/magic"

func isUpper(r rune) bool { return r >= 'A' && r <= 'Z' }

func isLower(r rune) bool { return r >= 'a' && r <= 'z' }

func isDigit(r rune) bool { return r >= '0' && r <= '9' }

func toUpper(r rune) rune {
	if isLower(r) {
		return r - ('a' - 'A')
	}
	return r
}

func toLower(r rune) rune {
	if isUpper(r) {
		return r + ('a' - 'A')
	}
	return r
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
		if isUpper(ch) || isLower(ch) || isDigit(ch) {
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
