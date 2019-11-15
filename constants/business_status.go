package constants

type ResponseStatus struct {
	Code   string
	Remark string
}

var (
	Exception  ResponseStatus = ResponseStatus{"Exception", "系统异常"}
	Success    ResponseStatus = ResponseStatus{"Success", "业务成功"}
	ParamError ResponseStatus = ResponseStatus{"ParamError", "参数异常"}
)
