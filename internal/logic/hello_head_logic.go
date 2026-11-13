// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"

	"example.com/hello-world/internal/svc"
	"example.com/hello-world/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloHeadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloHeadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloHeadLogic {
	return &HelloHeadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloHeadLogic) HelloHead() (*types.HelloResponse, error) {
	return NewHelloLogic(l.ctx, l.svcCtx).Hello()
}
