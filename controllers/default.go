package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"github.com/gocolly/colly"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	Temp()
	c.Data["Title"] = "图片库"
	c.Data["Content"] = "https://github.com/VampireMao/pic_web.git"
	c.TplName = "main/index.html"
}

func Temp() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("V", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(mahonia.NewDecoder("gbk").ConvertString(e.Attr("href")))
	})

	c.Visit("http://www.t66y.com/index.php")
}

func (c *MainController) Post() {
	println(c.Ctx.Request.URL.Path)
	println(c.Ctx.Request.Host)
}
