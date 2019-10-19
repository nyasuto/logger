package log

import "log"

type CustomLogger interface {
	Logger(format string, args ...interface{})
}

type customLog struct {
	output func(format string, args ...interface{})
}

func NewCustomLogger() *customLog {
	return &customLog{output: log.Printf}
}

func Logger(format string, args ...interface{}) {
	NewCustomLogger().Logger(format, args...)
}

// put prefix [ERROR] when the log argument contains go error
// by default putting prefix [INFO]
func (logger *customLog) Logger(format string, args ...interface{}) {
	prefix := "[INFO] "

	if args == nil {
		format = prefix + format
		log.Printf(format)
	} else {
		for _, arg := range args {
			switch arg.(type) {
			case error:
				prefix = "[ERROR] "
				break
			}
		}
		format = prefix + format
		log.Printf(format, args...)
	}
}
