package routers

import (
	"github.com/astaxie/beego"
	"pic_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/welcome", &controllers.WelcomeController{})
}
