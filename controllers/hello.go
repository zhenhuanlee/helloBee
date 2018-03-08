package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (hc *HelloController) Get() {
	hc.Data["hello"] = "hello world!"
	hc.TplName = "hello.tpl"
}
