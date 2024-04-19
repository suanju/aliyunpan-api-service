package auth

import (
	"context"

	"aliyunpan-api-service/internal/svc"
	"aliyunpan-api-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccessTokenFromRefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAccessTokenFromRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccessTokenFromRefreshTokenLogic {
	return &GetAccessTokenFromRefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccessTokenFromRefreshTokenLogic) GetAccessTokenFromRefreshToken(req *types.QrcodeStateReq) (resp *types.QrcodeStateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
