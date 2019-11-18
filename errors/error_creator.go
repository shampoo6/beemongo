package errors

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/shampoo6/beemongo/constants"
)

func CError(status constants.ResponseStatus, content interface{}) BusinessError {
	return BusinessError{status.Remark, status, content}
}

func CParamError(errors []*validation.Error) BusinessError {
	content := map[string]string{}
	for _, err := range errors {
		logs.Error(err.Key, err.Message)
		content[err.Key] = err.Message
	}
	return CError(constants.ParamError, content)
}
