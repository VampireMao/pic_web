package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Title"] = "图片库"
	c.Data["Content"] = "https://github.com/VampireMao/pic_web.git"
	c.TplName = "main/index.html"
}
