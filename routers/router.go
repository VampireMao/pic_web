package routers

import (
	"github.com/astaxie/beego"
	"pic_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// Admin
	beego.Router("/admin", &controllers.AdminController{}, "Get:Index")
	beego.Router("/admin/welcome", &controllers.AdminController{}, "Get:Welcome")
	// User
	beego.Router("/admin/user-list", &controllers.AdminController{}, "Get:UserListView")
	beego.Router("/admin/user-list", &controllers.AdminController{}, "Post:UserList")
	beego.Router("/admin/user-edit", &controllers.AdminController{}, "Post:UserEdit")
	beego.Router("/admin/user-delete", &controllers.AdminController{}, "Post:DeleteUser")
}
