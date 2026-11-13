// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"

	"example.com/hello-world/internal/svc"
	"example.com/hello-world/internal/types"
	"example.com/hello-world/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello() (*types.HelloResponse, error) {
	return &types.HelloResponse{
		Message: "Hello World!",
		User:    user.GetName(),
	}, nil
}
