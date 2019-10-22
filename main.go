package main

import (
	"github.com/astaxie/beego"
	_ "pic_web/routers"
)

func main() {
	beego.SetStaticPath("/public","static")
	beego.Run()
}

