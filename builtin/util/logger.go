package util

import (
	"go.uber.org/zap"
	"reflect"
)

var Logger *zap.Logger

func LoggerSymbols(logger *zap.Logger) map[string]reflect.Value {
	return map[string]reflect.Value{
		"Logger": reflect.ValueOf(logger),
	}
}
