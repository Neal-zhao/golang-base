package myLog

import "sync"

type MyLogger interface {
	Info(format string,args... interface{})
	Debug(format string,args... interface{})
	Close()
}
//等级变量
var mutex sync.Mutex
const MAX_SIZE = 1 << 5 * 1
type Level int64

const (
	INFO Level = iota
	WARNING
	DEBUG
	FATAL
)

func getLevel(level Level) string {
	switch level {
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case DEBUG:
		return "debug"
	case FATAL:
		return "fatal"
	default:
		return "debug"
	}
}
