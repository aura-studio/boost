package logger

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func New(s string) (*logrus.Logger, error) {
	logger := logrus.New()

	// Disable output.
	logger.SetOutput(io.Discard)

	// Enable caller.
	logger.SetReportCaller(true)

	// Set formatter.

	formatter, err := NewFormatter(gjson.Get(s, "fomatter").Raw)
	if err != nil {
		return nil, err
	}
	logger.SetFormatter(formatter)

	// Set hooks.
	gjson.Get(s, "hooks").ForEach(func(key, value gjson.Result) bool {
		typ := value.Get("type").String()
		hook, err := NewHook(typ, value.Raw)
		if err != nil {
			fmt.Printf("new hook error: %v", err)
			return false
		}
		logger.AddHook(hook)
		return true
	})

	// Set Level
	logger.SetLevel(NewLogLevel(gjson.Get(s, "level").Raw).ToLogrus())

	return logger, nil
}
