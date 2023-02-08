package systime

import (
	"github.com/aura-studio/boost/cast"
	"github.com/tidwall/gjson"
)

type SysTime struct {
	Zone  int64
	Fake  int64
	Delta int64
}

func Zone() int64 {
	return systime.Zone
}

func Fake() int64 {
	return systime.Fake
}

func Delta() int64 {
	return systime.Delta
}

var systime = &SysTime{}

func Init(s string) *SysTime {
	if result := gjson.Get(s, "zone"); result.Type != gjson.Null {
		WithTimeZone(result.String())
	} else {
		WithTimeZone("Asia/Shanghai")
	}

	if result := gjson.Get(s, "fake"); result.Type != gjson.Null {
		WithFakeTime(result.String())
	} else {
		WithFakeTime("0")
	}

	if result := gjson.Get(s, "delta"); result.Type != gjson.Null {
		WithDeltaTime(result.String())
	} else {
		WithDeltaTime("0")
	}

	return systime
}

type Options struct{}

var options = &Options{}

func WithTimeZone(tz string) {
	systime.Zone = options.parseTimeZone(tz)
}

func (*Options) parseTimeZone(s string) int64 {
	tz, err := cast.ToTimeZoneE(s)
	if err != nil {
		panic(err)
	}
	return cast.ToInt64(tz)
}

func WithFakeTime(s string) {
	systime.Fake = options.parseFakeTime(s)
}

func (*Options) parseFakeTime(s string) int64 {
	d, err := cast.ToDurationE(s)
	if err != nil {
		panic(err)
	}
	n, err := cast.ToInt64E(d)
	if err != nil {
		panic(err)
	}
	return n
}

func WithDeltaTime(s string) {
	systime.Delta = options.parseDeltaTime(s)
}

func (*Options) parseDeltaTime(s string) int64 {
	d, err := cast.ToDurationE(s)
	if err != nil {
		panic(err)
	}
	n, err := cast.ToInt64E(d)
	if err != nil {
		panic(err)
	}
	return n
}
