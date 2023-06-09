// Code generated by goctl. DO NOT EDIT.
// Source: rep.proto

package reper

import (
	"context"

	"YxEmr/sqd/rpc/rep/rep"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Req  = rep.Req
	Resp = rep.Resp

	Reper interface {
		Do(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
		Co(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	}

	defaultReper struct {
		cli zrpc.Client
	}
)

func NewReper(cli zrpc.Client) Reper {
	return &defaultReper{
		cli: cli,
	}
}

func (m *defaultReper) Do(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	client := rep.NewReperClient(m.cli.Conn())
	return client.Do(ctx, in, opts...)
}

func (m *defaultReper) Co(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	client := rep.NewReperClient(m.cli.Conn())
	return client.Co(ctx, in, opts...)
}
