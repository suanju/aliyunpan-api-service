package errorx

import (
	"aliyunpan-api-service/internal/utils/results"
	"github.com/pkg/errors"
	xerror "github.com/zeromicro/x/errors"
)

func NewCodeError(code int, err error) error {
	return &xerror.CodeMsg{
		Code: code,
		Msg:  err.Error(),
	}
}

func NewDefaultError(err error) error {
	return NewCodeError(results.CodeDefaultError, err)
}

func NewDefaultErrorMessage(msg string) error {
	return NewCodeError(results.CodeDefaultError, errors.New(msg))
}
