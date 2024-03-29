package models

import (
	"github.com/astaxie/beego/orm"
)

type UserModel struct {
	UserId   int    `orm:"pk;auto"`
	UserKey  string `orm:"size(64),unique"`
	UserName string `orm:"size(64)"`
	AuthStr  string `orm:"size(512)"`
	Password string `orm:"size(128)"`
	IsAdmin  int8   `orm:"default(0)"`
}

func (m *UserModel) TableName() string {
	return TbNameUser()
}

func (u *UserModel) TableUnique() [][]string {
	return [][]string{
		[]string{"UserKey"},
	}
}

func UserList(pageSize, page int) ([]*UserModel, int64) {
	offset := (page - 1) * pageSize

	query := orm.NewOrm().QueryTable(TbNameUser())
	total, _ := query.Count()

	data := make([]*UserModel, 0)
	query.Limit(pageSize, offset).All(&data, "UserId", "UserKey", "UserName")

	return data, total
}

func GetUserByName(username string) UserModel {
	user := UserModel{UserKey: username}
	o := orm.NewOrm()
	o.Read(&user, "user_key")
	return user
}

func UpdateUser(user *UserModel, cols string) error {
	o := orm.NewOrm()
	_, err := o.Update(user, cols)
	return err
}

func DeleteUser(user *UserModel) error {
	o := orm.NewOrm()
	_, err := o.Delete(user)
	return err
}
