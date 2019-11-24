package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

type WelcomeController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.Data["Title"] = "漫画屋 - 后台管理"
	c.TplName = "admin/index.html"
}

func (c *AdminController) Post() {
	println(c.Ctx.Request.URL.Path)
	println(c.Ctx.Request.Host)
}

func (c *WelcomeController) Get() {
	c.Data["Title"] = "欢迎页"
	c.TplName = "admin/welcome.html"
}
