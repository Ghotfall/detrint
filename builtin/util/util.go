package util

import (
	"go.uber.org/zap"
	"reflect"
)

var Logger *zap.Logger
var Vars map[string]interface{}

func Symbols(l *zap.Logger, v map[string]interface{}) map[string]reflect.Value {
	Logger = l
	Vars = v

	return map[string]reflect.Value{
		"Logger": reflect.ValueOf(Logger), // TODO: wrapper for logger?
		"Vars":   reflect.ValueOf(Vars),
	}
}
