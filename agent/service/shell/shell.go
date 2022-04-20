package shell

import (
	"bytes"
	"context"
	"github.com/ghotfall/detrint/grpc/shellpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os/exec"
	"strings"
)

var Logger *zap.Logger

func Register(s grpc.ServiceRegistrar, l *zap.Logger) {
	Logger = l

	shellpb.RegisterShellServiceServer(s, Server{})
	Logger.Info("Registered new service", zap.String("service", "shell"))
}

type Server struct {
	shellpb.UnimplementedShellServiceServer
}

func (s Server) Execute(_ context.Context, request *shellpb.ExecuteRequest) (*shellpb.ExecuteResponse, error) {
	Logger.Info("Method shell.execute is called", zap.String("request", request.String()))

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	resp := shellpb.ExecuteResponse{}

	cmd := exec.Command("powershell", "-NoProfile", request.GetScript())
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	resp.Stdout = strings.ToValidUTF8(stdout.String(), "")
	resp.Stderr = strings.ToValidUTF8(stderr.String(), "")

	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		if ok {
			resp.Code = int32(exitError.ExitCode())
		} else {
			Logger.Error("Unexpected error while executing shell.execute", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "failed to run script: %s", err.Error())
		}
	} else {
		resp.Code = 0
	}

	Logger.Debug(
		"Execution of shell.execute finished",
		zap.String("stdout", resp.Stdout),
		zap.String("stderr", resp.Stderr),
		zap.Int("code", int(resp.Code)),
	)

	return &resp, nil
}
