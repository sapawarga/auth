package grpc

import (
	"context"

	tpAuth "github.com/sapawarga/proto-file/auth"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	decodeToken   kitgrpc.Handler
	getUserDetail kitgrpc.Handler
}

func (g *grpcServer) DecodeToken(ctx context.Context, req *tpAuth.RequestDecodeToken) (*tpAuth.ResponsesDecodeToken, error) {
	_, resp, err := g.decodeToken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*tpAuth.ResponsesDecodeToken), nil
}

func (g *grpcServer) GetUserDetailByAccessToken(ctx context.Context, req *tpAuth.RequestDecodeToken) (*tpAuth.ResponseUserDetail, error) {
	_, resp, err := g.getUserDetail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*tpAuth.ResponseUserDetail), nil
}
