package types

type GenerateResponse struct {
	Content  GenerateContent `json:"content"`
	HasError bool            `json:"hasError"`
}

type GenerateContent struct {
	Data    GenerateData `json:"data"`
	Status  int          `json:"status"`
	Success bool         `json:"success"`
}

type GenerateData struct {
	T               int64  `json:"t"`
	CodeContent     string `json:"codeContent"`
	Ck              string `json:"ck"`
	ResultCode      int    `json:"resultCode"`
	ProcessFinished bool   `json:"processFinished"`
}

type QrcodeStateResponse struct {
	Content  QrcodeStateContent `json:"content"`
	HasError bool               `json:"hasError"`
}

type QrcodeStateContent struct {
	Data    QrcodeStateData `json:"data"`
	Status  int             `json:"status"`
	Success bool            `json:"success"`
}

type QrcodeStateData struct {
	LoginResult          string `json:"loginResult"`
	LoginSucResultAction string `json:"loginSucResultAction"`
	St                   string `json:"st"`
	QrCodeStatus         string `json:"qrCodeStatus"`
	LoginType            string `json:"loginType"`
	BizExt               string `json:"bizExt"`
	LoginScene           string `json:"loginScene"`
	ResultCode           int    `json:"resultCode"`
	AppEntrance          string `json:"appEntrance"`
	Smartlock            bool   `json:"smartlock"`
	ProcessFinished      bool   `json:"processFinished"`
}

type QrcodeStateBizExtData struct {
	PdsLoginResult QrcodeStatePdsLoginResult `json:"pds_login_result"`
}
type QrcodeStatePdsLoginResult struct {
	Role           string                    `json:"role"`
	UserData       QrcodeStateBizExtUserData `json:"userData"`
	LoginType      string                    `json:"loginType"`
	ExpiresIn      int                       `json:"expiresIn"`
	RequestID      string                    `json:"requestId"`
	State          string                    `json:"state"`
	IsFirstLogin   bool                      `json:"isFirstLogin"`
	NeedLink       bool                      `json:"needLink"`
	PathStatus     string                    `json:"pathStatus"`
	NickName       string                    `json:"nickName"`
	NeedRpVerify   bool                      `json:"needRpVerify"`
	Avatar         string                    `json:"avatar"`
	AccessToken    string                    `json:"accessToken"`
	UserName       string                    `json:"userName"`
	UserID         string                    `json:"userId"`
	DefaultDriveID string                    `json:"defaultDriveId"`
	ExistLink      []interface{}             `json:"existLink"`
	ExpireTime     string                    `json:"expireTime"`
	DataPinSetup   bool                      `json:"dataPinSetup"`
	TokenType      string                    `json:"tokenType"`
	DataPinSaved   bool                      `json:"dataPinSaved"`
	RefreshToken   string                    `json:"refreshToken"`
	Status         string                    `json:"status"`
}

type QrcodeStateBizExtUserData struct {
	DingDingRobotURL      string `json:"DingDingRobotUrl"`
	EncourageDesc         string `json:"EncourageDesc"`
	FeedBackSwitch        bool   `json:"FeedBackSwitch"`
	FollowingDesc         string `json:"FollowingDesc"`
	DingDingRobotURLSmall string `json:"ding_ding_robot_url"`
	EncourageDescSmall    string `json:"encourage_desc"`
	FeedBackSwitchSmall   bool   `json:"feed_back_switch"`
	FollowingDescSmall    string `json:"following_desc"`
}
