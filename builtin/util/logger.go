package util

import (
	"go.uber.org/zap"
	"reflect"
)

var Logger *zap.Logger

func LoggerSymbols(l *zap.Logger) map[string]reflect.Value {
	Logger = l

	return map[string]reflect.Value{
		"Logger": reflect.ValueOf(Logger),
	}
}
