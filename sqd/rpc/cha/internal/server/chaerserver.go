// Code generated by goctl. DO NOT EDIT.
// Source: cha.proto

package server

import (
	"context"

	"YxEmr/sqd/rpc/cha/cha"
	"YxEmr/sqd/rpc/cha/internal/logic"
	"YxEmr/sqd/rpc/cha/internal/svc"
)

type ChaerServer struct {
	svcCtx *svc.ServiceContext
	cha.UnimplementedChaerServer
}

func NewChaerServer(svcCtx *svc.ServiceContext) *ChaerServer {
	return &ChaerServer{
		svcCtx: svcCtx,
	}
}

func (s *ChaerServer) Do(ctx context.Context, in *cha.Req) (*cha.Resp, error) {
	l := logic.NewDoLogic(ctx, s.svcCtx)
	if err := l.Do(in); err != nil {
		l.Logger.Error(err)
		return &cha.Resp{
			Code: 0,
			Msg:  err.Error(),
			Data: "",
		}, nil
	} else {
		return &cha.Resp{
			Code: 1,
			Msg:  "ok",
			Data: "",
		}, nil
	}
}
