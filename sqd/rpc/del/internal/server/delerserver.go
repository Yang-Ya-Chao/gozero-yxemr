// Code generated by goctl. DO NOT EDIT.
// Source: del.proto

package server

import (
	"context"

	"YxEmr/sqd/rpc/del/del"
	"YxEmr/sqd/rpc/del/internal/logic"
	"YxEmr/sqd/rpc/del/internal/svc"
)

type DelerServer struct {
	svcCtx *svc.ServiceContext
	del.UnimplementedDelerServer
}

func NewDelerServer(svcCtx *svc.ServiceContext) *DelerServer {
	return &DelerServer{
		svcCtx: svcCtx,
	}
}

func (s *DelerServer) Do(ctx context.Context, in *del.Req) (*del.Resp, error) {
	l := logic.NewDoLogic(ctx, s.svcCtx)
	_,err := l.Do(in)
	if err != nil {
		l.Logger.Error(err)
	}
	return &del.Resp{},err
}
