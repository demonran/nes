package log

import (
	"github.com/labstack/gommon/log"
	"sync"
)

var (
	Logger *log.Logger
	once   sync.Once
)

const defaultFormat = `{"time":"${time_rfc3339}","level":"${level}","file":"${short_file}","line":"${line}"}`

func init() {
	Logger = log.New("")
	Logger.SetHeader(defaultFormat)
}
