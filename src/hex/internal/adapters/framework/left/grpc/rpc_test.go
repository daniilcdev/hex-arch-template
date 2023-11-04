package rpc

import (
	"context"
	"net"
	"testing"

	"github.com/daniilcdev/hex-arch-template/hex/internal/adapters/app/api"
	"github.com/daniilcdev/hex-arch-template/hex/internal/adapters/core/arithmetic"
	"github.com/daniilcdev/hex-arch-template/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/daniilcdev/hex-arch-template/hex/internal/adapters/framework/right/db"
	"github.com/daniilcdev/hex-arch-template/hex/internal/ports"

	"log"
	"os"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	// ports

	var dbaseAdapter ports.DbPort
	var arithmeticAdapter ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)

	if err != nil {
		log.Fatalf("failed to initiate db connection: %v", err)
	}

	arithmeticAdapter = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, arithmeticAdapter)

	gRPCAdapter = NewAdapter(appAdapter)
	pb.RegisterArithmeticServiceServer(grpcServer, gRPCAdapter)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func bufDialer(ctx context.Context, target string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}

	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 1, B: 1}
	answer, err := client.GetAddition(ctx, params)

	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(2), answer.Value)
}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 2, B: 1}
	answer, err := client.GetSubtraction(ctx, params)

	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(1), answer.Value)
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 2, B: 2}
	answer, err := client.GetMultiplication(ctx, params)

	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(4), answer.Value)
}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 4, B: 2}
	answer, err := client.GetDivision(ctx, params)

	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(2), answer.Value)
}
