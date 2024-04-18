package _const

const (
	AliyunDriveGenerate = "https://passport.aliyundrive.com/newlogin/qrcode/generate.do"
	AliyunQrcodeQuery   = "https://passport.aliyundrive.com/newlogin/qrcode/query.do"
)

const (
	QrCodeStatusNew       string = "NEW"
	QrCodeStatusScanned   string = "SCANED"
	QrCodeStatusExpired   string = "EXPIRED"
	QrCodeStatusCanceled  string = "CANCELED"
	QrCodeStatusConfirmed string = "CONFIRMED"
)

var QrCodeStatus = map[string]string{
	QrCodeStatusNew:       "请用阿里云盘 App 扫码",
	QrCodeStatusScanned:   "请在手机上确认",
	QrCodeStatusExpired:   "二维码已过期",
	QrCodeStatusCanceled:  "已取消",
	QrCodeStatusConfirmed: "已确认",
}
