package shell

import (
	"context"
	"github.com/ghotfall/detrint/grpc/shellpb"
	"google.golang.org/grpc"
	"reflect"
)

var client shellpb.ShellServiceClient

func Symbols(conn *grpc.ClientConn) map[string]reflect.Value {
	client = shellpb.NewShellServiceClient(conn)

	return map[string]reflect.Value{
		"Execute": reflect.ValueOf(Execute),
	}
}

func Execute(script string) (stdout, stderr string, code int, err error) {
	response, err := client.Execute(context.Background(), &shellpb.ExecuteRequest{Script: script})
	if err != nil {
		return "", "", 0, err
	} else {
		return response.GetStdout(), response.GetStderr(), int(response.GetCode()), nil
	}
}
