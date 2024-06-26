// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "aliyunpan-api-service/internal/handler/auth"
	"aliyunpan-api-service/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/generate_qrcode",
				Handler: auth.GenerateQrcodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/qrcode_state",
				Handler: auth.QrcodeStateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_access_token_from_refresh_token",
				Handler: auth.GetAccessTokenFromRefreshTokenHandler(serverCtx),
			},
		},
		rest.WithPrefix("/auth"),
	)
}
