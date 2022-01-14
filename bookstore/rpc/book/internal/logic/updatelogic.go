package logic

import (
	"context"

	"bookstore/rpc/book/book"
	"bookstore/rpc/book/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *book.AddReq) (*book.AddResp, error) {
	// todo: add your logic here and delete this line

	return &book.AddResp{}, nil
}
