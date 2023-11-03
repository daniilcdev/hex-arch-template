package rpc

import (
	"context"

	"hex-arch-template/internal/adapters/framework/left/grpc/pb"
)

func (a Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {

	answer, err := a.api.GetAddition(req.A, req.B)
	if err != nil {
		return &pb.Answer{}, err
	}
	return &pb.Answer{Value: answer}, nil
}

func (a Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {

	answer, err := a.api.GetAddition(req.A, req.B)
	if err != nil {
		return &pb.Answer{}, err
	}
	return &pb.Answer{Value: answer}, nil
}

func (a Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {

	answer, err := a.api.GetAddition(req.A, req.B)
	if err != nil {
		return &pb.Answer{}, err
	}
	return &pb.Answer{Value: answer}, nil
}

func (a Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {

	answer, err := a.api.GetAddition(req.A, req.B)
	if err != nil {
		return &pb.Answer{}, err
	}
	return &pb.Answer{Value: answer}, nil
}
