package style_test

import (
	"testing"

	"github.com/aura-studio/boost/style"
)

func Test_Standard(t *testing.T) {
	t.Log(style.Standardize("ab1_c2=d_e3f", "_"))
}

func TestUnstandard(t *testing.T) {
	t.Log(style.Unstandardize("m2A1b=Cd3Ef", "_"))
}
