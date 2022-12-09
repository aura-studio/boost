package radix_test

import (
	"testing"

	"github.com/aura-studio/boost/radix"
)

func TestRadix(t *testing.T) {
	c := radix.NewCharset([]byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_")).Shuffle(0)
	if c.Size() != 64 {
		t.Error("radix charset size error")
	}

	r := radix.New(8, c)

	if r.Decode(r.Encode(123456789)) != 123456789 {
		t.Error("radix decode error")
	}
}
