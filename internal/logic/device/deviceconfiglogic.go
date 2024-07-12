package device

import (
	"context"

	"sprinkler/internal/svc"
	"sprinkler/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceConfigLogic {
	return &DeviceConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceConfigLogic) DeviceConfig(req *types.DeviceConfigReq) (resp *types.DeviceConfigResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.DeviceConfigResp{}
	resp.DeviceConfigId = "1"
	resp.DeviceConfigDetails = "test"
	return resp, nil
}
