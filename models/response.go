package models

import (
	"github.com/shampoo6/beemongo/constants"
	"github.com/shampoo6/beemongo/errors"
)

type Response struct {
	msg     string
	status  string
	content interface{}
}

// 创建成功响应, controller调用
func CSuccessResponse(content interface{}) map[string]interface{} {
	return map[string]interface{}{
		"msg":     constants.Success.Remark,
		"status":  constants.Success.Code,
		"content": content,
	}
}

// 创建系统异常响应=
func CExceptionResponse(content string) map[string]interface{} {
	return map[string]interface{}{
		"msg":     constants.Exception.Remark,
		"status":  constants.Exception.Code,
		"content": content,
	}
}

func CBusinessResponse(err errors.BusinessError) map[string]interface{} {
	return map[string]interface{}{
		"msg":     err.Msg,
		"status":  err.Status.Code,
		"content": err.Content,
	}
}
