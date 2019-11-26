package constants

type ResponseStatus struct {
	Code   string
	Remark string
}

var (
	Exception  = ResponseStatus{"Exception", "系统异常"}
	AuthError  = ResponseStatus{"AuthError", "权限错误"}
	Success    = ResponseStatus{"Success", "业务成功"}
	ParamError = ResponseStatus{"ParamError", "参数错误"}
)
