package style_test

import (
	"testing"

	"github.com/aura-studio/boost/style"
)

func Test_Standard(t *testing.T) {
	t.Log(style.Standardize("ab_c=d_ef", "_"))
}

func TestUnstandard(t *testing.T) {
	t.Log(style.Unstandardize("m2A1b=Cd3Ef", "_"))
}
