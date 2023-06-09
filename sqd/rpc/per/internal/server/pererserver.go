// Code generated by goctl. DO NOT EDIT.
// Source: per.proto

package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"YxEmr/sqd/rpc/per/internal/logic"
	"YxEmr/sqd/rpc/per/internal/svc"
	"YxEmr/sqd/rpc/per/per"
)

type PererServer struct {
	svcCtx *svc.ServiceContext
	per.UnimplementedPererServer
}

func NewPererServer(svcCtx *svc.ServiceContext) *PererServer {
	return &PererServer{
		svcCtx: svcCtx,
	}
}

func (s *PererServer) Do(ctx context.Context, in *per.Req) (*per.Resp, error) {
	l := logic.NewDoLogic(ctx, s.svcCtx)
	r,err := l.Do(in)
	if err != nil {
		l.Logger.Error(err)
		err = status.Error(codes.Aborted, err.Error())
	}
	return r,err
}

func (s *PererServer) Co(ctx context.Context, in *per.Req) (*per.Resp, error) {
	l := logic.NewCoLogic(ctx, s.svcCtx)
	r,err := l.Co(in)
	if err != nil {
		l.Logger.Error(err)
		err = status.Error(codes.Aborted, err.Error())
	}
	return r,err
}
