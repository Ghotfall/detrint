package file

import (
	"context"
	"github.com/ghotfall/detrint/grpc/filepb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

var logger *zap.Logger

func Register(s grpc.ServiceRegistrar, l *zap.Logger) {
	logger = l

	filepb.RegisterFileServiceServer(s, Server{})
	logger.Info("Registered new service", zap.String("service", "file"))
}

type Server struct {
	filepb.UnimplementedFileServiceServer
}

func (s Server) GetStat(_ context.Context, request *filepb.StatRequest) (*filepb.StatResponse, error) {
	logger.Info("Method file.getstat is called", zap.String("request", request.String()))

	filename := request.GetFilename()
	resp := filepb.StatResponse{}

	fileInfo, err := os.Stat(filename)
	if err != nil {
		pathErr, ok := err.(*os.PathError)
		if ok {
			resp.PathErr = pathErr.Error()
		} else {
			return nil, status.Errorf(codes.Internal, "failed to get file stat: %s", err.Error())
		}
	} else {
		resp.Name = fileInfo.Name()
		resp.Size = fileInfo.Size()
		resp.Mode = uint32(fileInfo.Mode())
		resp.Time = fileInfo.ModTime().Format(time.RFC3339)
		resp.Dir = fileInfo.IsDir()
	}

	logger.Debug("Execution of shell.execute finished", zap.Any("result", &resp))
	return &resp, nil
}
