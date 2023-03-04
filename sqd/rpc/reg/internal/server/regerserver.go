// Code generated by goctl. DO NOT EDIT.
// Source: reg.proto

package server

import (
	"context"

	"YxEmr/sqd/rpc/reg/internal/logic"
	"YxEmr/sqd/rpc/reg/internal/svc"
	"YxEmr/sqd/rpc/reg/reg"
)

type RegerServer struct {
	svcCtx *svc.ServiceContext
	reg.UnimplementedRegerServer
}

func NewRegerServer(svcCtx *svc.ServiceContext) *RegerServer {
	return &RegerServer{
		svcCtx: svcCtx,
	}
}

func (s *RegerServer) Do(ctx context.Context, in *reg.Req) (*reg.Resp, error) {
	l := logic.NewDoLogic(ctx, s.svcCtx)
	if err := l.Do(in); err != nil {
		l.Logger.Error(err)
		return &reg.Resp{
			Code: 0,
			Msg:  err.Error(),
			Data: "",
		}, nil
	} else {
		return &reg.Resp{
			Code: 1,
			Msg:  "ok",
			Data: "",
		}, nil
	}
}
