package config

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/aura-studio/boost/cast"
)

func Assert(condition bool, v interface{}, args ...string) {
	k := strings.Join(args, ".")
	if !condition {
		panic(fmt.Errorf("invalid config value [%v] for key [%v] ", v, k))
	}
}

func JSON(args ...string) string {
	if len(args) == 0 {
		data, err := json.Marshal(c.AllSettings())
		if err != nil {
			panic(err)
		}
		return string(data)
	}
	s := strings.Join(args, ".")
	data, err := json.Marshal(c.Sub(s).AllSettings())
	if err != nil {
		panic(err)
	}
	return string(data)
}

// regexp to match ${} or $()
var re = regexp.MustCompile(`\$\{([^}]+)\}|\$\(([^\)]+)\)`)

func Parse(s string) string {
	return re.ReplaceAllStringFunc(s, func(v string) string {
		k := v[2 : len(v)-1]
		return cast.ToString(Get(k))
	})
}
