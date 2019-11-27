package models

import (
	//"strconv"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(UserModel), new(ComicCategoryModel))
}

func TbNameUser() string {
	return "xcms_user"
}

func TbNameComicCategory() string {
	return "xcms_comic_category"
}
