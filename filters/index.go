package filters

import "github.com/astaxie/beego"

func InitFilters() error {
	beego.InsertFilter("*", beego.BeforeRouter, mainFilter)
	return nil
}
