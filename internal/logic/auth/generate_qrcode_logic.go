package auth

import (
	_const "aliyunpan-api-service/internal/const"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/skip2/go-qrcode"

	"aliyunpan-api-service/internal/svc"
	"aliyunpan-api-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateQrcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateQrcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateQrcodeLogic {
	return &GenerateQrcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateQrcodeLogic) GenerateQrcode(req *types.GenerateQrcodeReq) (resp *types.GenerateQrcodeResp, err error) {
	clientResp, err := resty.New().R().
		SetQueryParams(map[string]string{
			"appName":     "aliyun_drive",
			"appEntrance": "web",
			"isMobile":    "false",
		}).
		Get(_const.AliyunDriveGenerate)
	if err != err {
		return nil, err
	}
	httpResp := new(types.GenerateResponse)
	err = json.Unmarshal(clientResp.Body(), httpResp)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err %s", clientResp.Body())
	}
	//生成二维码
	qrImg, err := qrcode.Encode(httpResp.Content.Data.CodeContent, qrcode.Medium, 256)
	if err != nil {
		return nil, fmt.Errorf("failed to generate qrcode %s", err)
	}
	base64Str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(qrImg)
	return &types.GenerateQrcodeResp{
		Qrcode: base64Str,
		T:      httpResp.Content.Data.T,
		Ck:     httpResp.Content.Data.Ck,
	}, nil
}
