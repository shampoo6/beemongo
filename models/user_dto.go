package models

import (
	"github.com/astaxie/beego/validation"
	"regexp"
	"strings"
)

type UserDto struct {
	Id string
	// 使用 valid 命令的方式验证参数时，无法修改返回的错误提示
	//Name string `valid:"Required"`
	//Sex  string `valid:"Required;Match(/^(Male|Female)$/)"`
	//Age  int    `valid:"Required;Range(0, 200)"`
	Name string
	Sex  string
	Age  int
}

func (u *UserDto) Validation() *validation.Validation {
	ptr, _ := regexp.Compile("^(Male|Female)$")
	valid := validation.Validation{}
	valid.Required(u.Name, "Name").Message("姓名不能为空")
	valid.Match(u.Sex, ptr, "Sex").Message("性别必须为 Male 或 Female")
	valid.Required(u.Age, "Age").Message("年龄不能为空")
	valid.Range(u.Age, 0, 200, "Age").Message("年龄必须在0~200之间")
	return &valid
}

func (u *UserDto) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}
