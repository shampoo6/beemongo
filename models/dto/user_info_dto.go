package dto

import (
	"github.com/astaxie/beego/validation"
	"github.com/shampoo6/beemongo/utils/string_util"
)

type UserInfoDto struct {
	Id      string
	UserId  string
	Email   string
	Mobile  string
	Address string
	Name    string
}

func (u *UserInfoDto) Validation() *validation.Validation {
	valid := validation.Validation{}
	valid.Required(u.UserId, "UserId").Message("用户id不能为空")
	if string_util.HasText(u.Email) {
		valid.Email(u.Email, "Email").Message("邮件格式不正确")
	}
	if string_util.HasText(u.Mobile) {
		valid.Mobile(u.Mobile, "Mobile").Message("手机号格式不正确")
	}
	return &valid
}
