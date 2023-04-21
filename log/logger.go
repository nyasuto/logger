package log

import "log"

type CustomLogger interface {
	Logger(format string, args ...any)
}

type customLog struct {
	output func(format string, args ...any)
}

func NewCustomLogger() *customLog {
	return &customLog{output: log.Printf}
}

func Logger(format string, args ...any) {
	NewCustomLogger().Logger(format, args...)
}

// put prefix [ERROR] when the log argument contains go error
// by default putting prefix [INFO]
func (logger *customLog) Logger(format string, args ...any) {
	prefix := "[INFO] "

	if args == nil {
		format = prefix + format
		logger.output(format)
	} else {
		for _, arg := range args {
			switch arg.(type) {
			case error:
				prefix = "[ERROR] "				
			}
		}
		format = prefix + format
		logger.output(format, args...)
	}
}
