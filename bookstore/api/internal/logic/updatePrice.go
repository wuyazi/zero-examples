package logic

import (
	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/rpc/book/book"
	"context"
	"github.com/tal-tech/go-zero/core/logx"
)

type UpdatedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdatedLogic {
	return UpdatedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatedLogic) UpdatePrice(req types.UpdateReq) (*types.AddResp, error) {
	resp, err := l.svcCtx.BookClient.UpdatePrice(l.ctx, &book.UpdatePriceReq{
		Id:    req.Id,
		Price: req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddResp{
		Id: resp.Id,
	}, nil
}
