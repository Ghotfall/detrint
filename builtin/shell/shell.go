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

func Execute(script string) (string, error) {
	response, err := client.Execute(context.Background(), &shellpb.ExecuteRequest{Script: script})
	if err != nil {
		return "", err
	} else {
		return response.GetResult(), nil
	}
}
