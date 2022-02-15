// Code generated by goctl. DO NOT EDIT!
// Source: pay.proto

package payclient

import (
	"context"

	"mall/service/pay/rpc/pay"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	CallbackResponse = pay.CallbackResponse
	CreateRequest    = pay.CreateRequest
	CreateResponse   = pay.CreateResponse
	DetailRequest    = pay.DetailRequest
	DetailResponse   = pay.DetailResponse
	CallbackRequest  = pay.CallbackRequest

	Pay interface {
		Create(ctx context.Context, in *CreateRequest) (*CreateResponse, error)
		Detail(ctx context.Context, in *DetailRequest) (*DetailResponse, error)
		Callback(ctx context.Context, in *CallbackRequest) (*CallbackResponse, error)
	}

	defaultPay struct {
		cli zrpc.Client
	}
)

func NewPay(cli zrpc.Client) Pay {
	return &defaultPay{
		cli: cli,
	}
}

func (m *defaultPay) Create(ctx context.Context, in *CreateRequest) (*CreateResponse, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.Create(ctx, in)
}

func (m *defaultPay) Detail(ctx context.Context, in *DetailRequest) (*DetailResponse, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.Detail(ctx, in)
}

func (m *defaultPay) Callback(ctx context.Context, in *CallbackRequest) (*CallbackResponse, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.Callback(ctx, in)
}
