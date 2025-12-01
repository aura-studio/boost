package style

import (
	"strings"

	"github.com/aura-studio/boost/magic"
)

var (
	googleChain = *NewChainStyle(magic.SeparatorSlash, magic.SeparatorHyphen)
	unixChain   = *NewChainStyle(magic.SeparatorPeriod, magic.SeparatorUnderscore)
)

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
