package rpc

import (
	context "context"
	"github.com/projectxpolaris/youlog/config"
	"github.com/projectxpolaris/youlog/pb"
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
	pb.RegisterLogServiceServer(rpcServer, &l.server)
	log.Printf("server listening at %v", lis.Addr())
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Server struct {
	pb.UnimplementedLogServiceServer
}

func (l *Server) WriteLog(ctx context.Context, data *pb.LogData) (*pb.WriteReply, error) {
	service.DefaultLogWriter.In <- data
	return &pb.WriteReply{Success: true}, nil
}
