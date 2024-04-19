package auth

import (
	"aliyunpan-api-service/internal/utils/errorx"
	"aliyunpan-api-service/internal/utils/results"
	"net/http"

	"aliyunpan-api-service/internal/logic/auth"
	"aliyunpan-api-service/internal/svc"
	"aliyunpan-api-service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	zeroMicroHttp "github.com/zeromicro/x/http"
)

func QrcodeStateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QrcodeStateReq
		if err := httpx.Parse(r, &req); err != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeError(results.CodeTypeError, err))
			return
		}
		if validateErr := svcCtx.Validate(&req); validateErr != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeError(results.CodeTypeError, validateErr))
			return
		}

		l := auth.NewQrcodeStateLogic(r.Context(), svcCtx)
		resp, err := l.QrcodeState(&req)
		if err != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
