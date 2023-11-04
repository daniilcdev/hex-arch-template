package rpc

import (
	"log"
	"net"

	"github.com/daniilcdev/hex-arch-template/src/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/daniilcdev/hex-arch-template/src/hex/internal/ports"

	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServerService := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServerService)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server gRPC server over port 9000: %v", err)
	}
}
