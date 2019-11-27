package main

import (
	"github.com/astaxie/beego"
	_ "pic_web/routers"
	_ "pic_web/sysinit"
)

func main() {
	beego.SetStaticPath("/public", "static")
	beego.Run()
}
