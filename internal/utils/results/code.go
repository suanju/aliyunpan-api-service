package results

const (
	CodeSuccess int = 0

	CodeTypeError    int = 415
	CodeDefaultError int = 1001
)

var _ = map[int]string{
	CodeSuccess: "请求成功",

	CodeTypeError:    "类型错误",
	CodeDefaultError: "默认错误返回",
}
