package rpc

import (
	context "context"
	"github.com/project-xpolaris/youplustoolkit/youlog/logservice"
	"github.com/projectxpolaris/youlog/config"
	"github.com/projectxpolaris/youlog/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

var DefaultLogServer = &LogServer{}

type LogServer struct {
	server Server
}

func (l *LogServer) Run() {
	lis, err := net.Listen("tcp", config.Instance.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcServer := grpc.NewServer()
	l.server = Server{}
	logservice.RegisterLogServiceServer(rpcServer, &l.server)
	log.Printf("server listening at %v", lis.Addr())
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Server struct {
	logservice.UnimplementedLogServiceServer
}

func (l *Server) WriteLog(ctx context.Context, data *logservice.LogData) (*logservice.WriteReply, error) {
	service.DefaultLogWriter.In <- data
	return &logservice.WriteReply{Success: true}, nil
}
