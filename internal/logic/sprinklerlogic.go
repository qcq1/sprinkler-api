package logic

import (
	"context"
	"sprinkler/internal/svc"
	"sprinkler/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SprinklerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSprinklerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SprinklerLogic {
	return &SprinklerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SprinklerLogic) Sprinkler(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
