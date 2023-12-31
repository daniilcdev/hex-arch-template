package rpc

import (
	"context"

	"github.com/daniilcdev/hex-arch-template/src/hex/internal/adapters/framework/left/grpc/pb"
)

func (a Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {

	answer, err := a.api.GetAddition(req.A, req.B)
	if err != nil {
		return &pb.Answer{}, err
	}
	return &pb.Answer{Value: answer}, nil
}

func (a Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {

	answer, err := a.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return &pb.Answer{}, err
	}
	return &pb.Answer{Value: answer}, nil
}

func (a Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {

	answer, err := a.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return &pb.Answer{}, err
	}
	return &pb.Answer{Value: answer}, nil
}

func (a Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {

	answer, err := a.api.GetDivision(req.A, req.B)
	if err != nil {
		return &pb.Answer{}, err
	}
	return &pb.Answer{Value: answer}, nil
}
