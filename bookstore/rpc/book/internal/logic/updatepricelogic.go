package logic

import (
	"context"

	"bookstore/rpc/book/book"
	"bookstore/rpc/book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdatePriceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePriceLogic {
	return &UpdatePriceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePriceLogic) UpdatePrice(in *book.UpdatePriceReq) (*book.AddResp, error) {
	// todo: add your logic here and delete this line

	return &book.AddResp{}, nil
}
