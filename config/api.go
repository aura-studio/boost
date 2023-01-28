package config

import (
	"fmt"
	"strings"
)

func Assert(condition bool, v interface{}, args ...string) {
	k := strings.Join(args, ".")
	if !condition {
		panic(fmt.Errorf("invalid config value [%v] for key [%v] ", v, k))
	}
}
