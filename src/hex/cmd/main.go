package main

import (
	"log"
	"os"

	"github.com/daniilcdev/hex-arch-template/hex/internal/adapters/app/api"
	"github.com/daniilcdev/hex-arch-template/hex/internal/adapters/core/arithmetic"
	rpc "github.com/daniilcdev/hex-arch-template/hex/internal/adapters/framework/left/grpc"
	"github.com/daniilcdev/hex-arch-template/hex/internal/adapters/framework/right/db"
	"github.com/daniilcdev/hex-arch-template/hex/internal/ports"
)

func main() {
	var err error

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

	defer dbaseAdapter.CloseDbConnection()

	arithmeticAdapter = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, arithmeticAdapter)

	gRPCAdapter = rpc.NewAdapter(appAdapter)

	gRPCAdapter.Run()
}
