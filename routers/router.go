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
	beego.Router("/admin/add-user", &controllers.AdminController{}, "Get:AddUserView")
	beego.Router("/admin/add-user", &controllers.AdminController{}, "Post:AddUser")

	// ComicCate
	beego.Router("/admin/comic_cate-list", &controllers.AdminController{}, "Get:ComicCateListView")
	beego.Router("/admin/comic_cate-list", &controllers.AdminController{}, "Post:ComicCateList")
	beego.Router("/admin/comic_cate-edit", &controllers.AdminController{}, "Post:ComicCateEdit")
}
