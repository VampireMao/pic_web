package controllers

import (
	"github.com/astaxie/beego"
	"pic_web/models"
	"xcms/consts"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Index() {
	c.Data["Title"] = "漫画屋 - 后台管理"
	c.TplName = "admin/index.html"
}

func (c *AdminController) UserListView() {
	c.Data["Title"] = "欢迎页"
	c.TplName = "admin/user/user-list.html"
}

func (c *AdminController) UserList() {
	page, _ := c.GetInt("page", 1)
	limit, _ := c.GetInt("limit", 5)
	res, count := models.UserList(limit, page)
	r := &models.ListJsonResult{consts.JRCodeSucc, "OK", count, res}
	c.Data["json"] = r
	c.ServeJSON()
}

func (c *AdminController) Welcome() {
	c.Data["Title"] = "欢迎页"
	c.TplName = "admin/welcome.html"
}

func (c *AdminController) UserEdit() {
	userId, _ := c.GetInt("UserId", 0)
	cols := c.GetString("cols")
	u := models.UserModel{
		UserId:   userId,
		UserName: c.GetString("UserName"),
	}
	err := models.UpdateUser(&u, cols)
	var r *models.JsonResult
	if err != nil {
		r = &models.JsonResult{
			Code: consts.JRCodeFailed,
			Msg:  "Error",
			Data: err,
		}
	} else {
		r = &models.JsonResult{
			Code: consts.JRCodeSucc,
			Msg:  "OK",
			Data: nil,
		}
	}
	c.Data["json"] = r
	c.ServeJSON()
}

func (c *AdminController) DeleteUser() {
	userId, _ := c.GetInt("UserId")
	u := models.UserModel{
		UserId: userId,
	}
	err := models.DeleteUser(&u)
	var r *models.JsonResult
	if err != nil {
		r = &models.JsonResult{
			Code: consts.JRCodeFailed,
			Msg:  "Error",
			Data: err,
		}
	} else {
		r = &models.JsonResult{
			Code: consts.JRCodeSucc,
			Msg:  "OK",
			Data: nil,
		}
	}
	c.Data["json"] = r
	c.ServeJSON()
}

func (c *AdminController) ComicCateListView() {
	c.Data["Title"] = "欢迎页"
	c.TplName = "admin/comic/comic_cate-list.html"
}

func (c *AdminController) ComicCateList() {
	page, _ := c.GetInt("page", 1)
	limit, _ := c.GetInt("limit", 5)
	res, count := models.ComicCateList(limit, page)
	r := &models.ListJsonResult{consts.JRCodeSucc, "OK", count, res}
	c.Data["json"] = r
	c.ServeJSON()
}

func (c *AdminController) ComicCateEdit() {
	id, _ := c.GetInt("id", 0)
	zhName := c.GetString("zhName")
	enName := c.GetString("enName")
	isAdult, _ := c.GetInt8("isAdult", 0)
	cols := c.GetString("cols")
	cc := models.ComicCategoryModel{
		Id:      id,
		ZhName:  zhName,
		EnName:  enName,
		IsAdult: isAdult,
	}

	err := models.UpdateComicCate(&cc, cols)
	var r *models.JsonResult
	if err != nil {
		r = &models.JsonResult{
			Code: consts.JRCodeFailed,
			Msg:  "Error",
			Data: err,
		}
	} else {
		r = &models.JsonResult{
			Code: consts.JRCodeSucc,
			Msg:  "OK",
			Data: nil,
		}
	}

	c.Data["json"] = r
	c.ServeJSON()
}

func (c *AdminController) AddUserView() {
	c.Data["Title"] = "新增用户"
	c.TplName = "admin/user/admin-add.html"
}

func (c *AdminController) AddUser() {

}
