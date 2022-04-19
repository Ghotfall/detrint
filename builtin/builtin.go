package builtin

import (
	"github.com/ghotfall/detrint/builtin/util"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"go.uber.org/zap"
	"reflect"
)

func Load(interpreter *interp.Interpreter, l *zap.Logger) error {
	stdlibErr := interpreter.Use(stdlib.Symbols)
	if stdlibErr != nil {
		return stdlibErr
	}

	// github.com/ghotfall/detrint/builtin
	symbols := make(map[string]map[string]reflect.Value)
	symbols["github.com/ghotfall/detrint/builtin/util/util"] = util.LoggerSymbols(l)
	symbolsErr := interpreter.Use(symbols)
	if symbolsErr != nil {
		return symbolsErr
	}

	return nil
}
