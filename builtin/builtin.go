package builtin

import (
	"github.com/ghotfall/detrint/builtin/util"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"reflect"
)

type Settings struct {
	Logger     *zap.Logger
	Vars       map[string]interface{}
	Connection *grpc.ClientConn
}

func Load(interpreter *interp.Interpreter, settings Settings) error {
	stdlibErr := interpreter.Use(stdlib.Symbols)
	if stdlibErr != nil {
		return stdlibErr
	}

	symbols := make(map[string]map[string]reflect.Value)

	symbols["github.com/ghotfall/detrint/builtin/util/util"] = util.LoggerSymbols(settings.Logger, settings.Vars)

	symbolsErr := interpreter.Use(symbols)
	if symbolsErr != nil {
		return symbolsErr
	}

	return nil
}
