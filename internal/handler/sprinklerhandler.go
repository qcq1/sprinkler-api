package handler

import (
	"net/http"
	"sprinkler/internal/logic"
	"sprinkler/internal/svc"
	"sprinkler/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SprinklerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSprinklerLogic(r.Context(), svcCtx)
		resp, err := l.Sprinkler(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
