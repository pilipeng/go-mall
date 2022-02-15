package logic

import (
	"context"
	"mall/service/product/rpc/product"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req types.UpdateRequest) (*types.UpdateResponse, error) {
	_, err := l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:     req.Id,
		Name:   req.Name,
		Desc:   req.Desc,
		Stock:  req.Stock,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}