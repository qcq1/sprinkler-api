package device

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sprinkler/internal/logic/device"
	"sprinkler/internal/svc"
	"sprinkler/internal/types"
)

func DeviceConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceConfigReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := device.NewDeviceConfigLogic(r.Context(), svcCtx)
		resp, err := l.DeviceConfig(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
