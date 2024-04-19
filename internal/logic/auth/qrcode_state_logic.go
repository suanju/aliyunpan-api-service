package auth

import (
	_const "aliyunpan-api-service/internal/const"
	"aliyunpan-api-service/internal/svc"
	"aliyunpan-api-service/internal/types"
	"aliyunpan-api-service/internal/utils/errorx"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
	"strconv"
)

type QrcodeStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQrcodeStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QrcodeStateLogic {
	return &QrcodeStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QrcodeStateLogic) QrcodeState(req *types.QrcodeStateReq) (resp *types.QrcodeStateResp, err error) {
	params := url.Values{}
	params.Set("t", strconv.FormatInt(req.T, 10))
	params.Set("ck", req.Ck)
	params.Set("appName", "aliyun_drive")
	params.Set("appEntrance", "web")
	body := params.Encode()
	clientResp, err := resty.New().R().
		SetHeader(_const.HeaderContentType, _const.FormURLEncodedContentType).
		SetBody(body).
		Post(_const.AliyunQrcodeQuery)

	httpResp := new(types.QrcodeStateResponse)
	err = json.Unmarshal(clientResp.Body(), httpResp)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err %s", clientResp.Body())
	}
	if httpResp.Content.Data.QrCodeStatus == _const.QrCodeStatusConfirmed {
		//成功扫码授权
		decodedBytes, _ := base64.StdEncoding.DecodeString(httpResp.Content.Data.BizExt)
		bizExt := new(types.QrcodeStateBizExtData)
		_ = json.Unmarshal(decodedBytes, &bizExt)
		return &types.QrcodeStateResp{
			UserName:     bizExt.PdsLoginResult.UserName,
			UserID:       bizExt.PdsLoginResult.UserID,
			TokenType:    bizExt.PdsLoginResult.TokenType,
			RefreshToken: bizExt.PdsLoginResult.RefreshToken,
		}, nil
	}
	return nil, errorx.NewDefaultErrorMessage(_const.QrCodeStatus[httpResp.Content.Data.QrCodeStatus])
}
