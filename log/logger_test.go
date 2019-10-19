package log

import (
	"errors"
	"fmt"
	"gotest.tools/assert"
	"testing"
)

// Logger test verify these case do not cause a panic
// I have checked PREFIX by eyeballs

var formatResult = ""

func dummyPrintf(format string, args ...interface{}) {
	formatResult = fmt.Sprintf(format, args...)
}

func TestCustomLoggerWhenNoArgument(t *testing.T) {

	logger := customLog{output: dummyPrintf}
	logger.Logger("this is text only log\n")
	assert.Equal(t, "[INFO] this is text only log\n", formatResult)
}

func TestCustomLoggerWithError(t *testing.T) {
	logger := customLog{output: dummyPrintf}
	logger.Logger("%v\n", errors.New("this is error log"))
	assert.Equal(t, "[ERROR] this is error log\n", formatResult)
}

func TestCustomLoggerWithInfo(t *testing.T) {
	logger := customLog{output: dummyPrintf}
	logger.Logger("this is information plus number %v\n", 10)
	assert.Equal(t, "[INFO] this is information plus number 10\n", formatResult)
}

func TestCustomLoggerWithInfoAndError(t *testing.T) {
	logger := customLog{output: dummyPrintf}
	logger.Logger("number plus error should be error %d %s\n", 10, errors.New("this is error"))
	assert.Equal(t, "[ERROR] number plus error should be error 10 this is error\n", formatResult)
}
