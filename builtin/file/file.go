package file

import (
	"context"
	"github.com/ghotfall/detrint/grpc/filepb"
	"google.golang.org/grpc"
	"reflect"
)

var client filepb.FileServiceClient

func Symbols(conn *grpc.ClientConn) map[string]reflect.Value {
	client = filepb.NewFileServiceClient(conn)

	return map[string]reflect.Value{
		"GetStat": reflect.ValueOf(GetStat),
		"Info":    reflect.ValueOf((*Info)(nil)),
	}
}

type Info struct {
	Name    string
	Size    int64
	Mode    uint32
	Time    string
	Dir     bool
	PathErr string
}

func GetStat(filename string) (Info, error) {
	stat, err := client.GetStat(context.Background(), &filepb.StatRequest{Filename: filename})
	if err != nil {
		return Info{}, err
	} else {
		return Info{
			Name:    stat.GetName(),
			Size:    stat.GetSize(),
			Mode:    stat.GetMode(),
			Time:    stat.GetTime(),
			Dir:     stat.GetDir(),
			PathErr: stat.GetPathErr(),
		}, nil
	}
}
