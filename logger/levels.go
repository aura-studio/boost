package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

var (
	ErrUnknownLogLevel = errors.New("unknown log level")
	ErrLevelNotSet     = errors.New("log level not set")
)

type LogLevel struct {
	level logrus.Level
}

func NewLogLevel(s string) *LogLevel {
	if json.Valid([]byte(s)) {
		s = gjson.Get(s, "level").String()
	}

	var level logrus.Level

	switch strings.ToLower(s) {
	case "panic":
		level = logrus.PanicLevel
	case "fatal":
		level = logrus.FatalLevel
	case "error":
		level = logrus.ErrorLevel
	case "warn", "warning":
		level = logrus.WarnLevel
	case "info", "print":
		level = logrus.InfoLevel
	case "debug":
		level = logrus.DebugLevel
	case "trace":
		level = logrus.TraceLevel
	default:
		panic(fmt.Errorf("%w: %s", ErrUnknownLogLevel, s))
	}

	return &LogLevel{level}
}

func (l *LogLevel) ToLogrus() logrus.Level {
	return l.level
}

type LogLevels struct {
	logLevels []*LogLevel
}

func NewLogLevels(i interface{}) *LogLevels {
	switch i := i.(type) {
	case string:
		s := i

		if json.Valid([]byte(s)) {
			if gjson.Get(s, "level").IsArray() {
				s = gjson.Get(s, "level").Raw
			} else {
				s = gjson.Get(s, "level").String()
			}
		}

		logLevel := NewLogLevel(s)
		logLevels := make([]*LogLevel, 0, logLevel.ToLogrus()+1)

		for i := logrus.Level(0); i <= logLevel.ToLogrus(); i++ {
			s := logrus.Level(i).String()
			logLevels = append(logLevels, NewLogLevel(s))
		}

		return &LogLevels{logLevels}

	case []string:
		strs := i

		logLevels := make([]*LogLevel, 0, len(strs))
		for _, str := range strs {
			logLevels = append(logLevels, NewLogLevel(str))
		}

		return &LogLevels{logLevels}
	default:
		panic(fmt.Errorf("%w: %T", ErrUnknownLogLevel, i))
	}
}

func (l *LogLevels) ToLogrus() []logrus.Level {
	logrusLevels := make([]logrus.Level, 0, len(l.logLevels))
	for _, level := range l.logLevels {
		logrusLevels = append(logrusLevels, level.ToLogrus())
	}

	return logrusLevels
}