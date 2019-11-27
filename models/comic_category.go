package models

import "github.com/astaxie/beego/orm"

type ComicCategoryModel struct {
	Id      int    `orm:"pk;auto"`
	ZhName  string `orm:"size(64),unique"`
	EnName  string `orm:"size(64)"`
	IsAdult int8   `orm:"default(0)"`
}

func (m *ComicCategoryModel) TableName() string {
	return TbNameComicCategory()
}

func (u *ComicCategoryModel) TableUnique() [][]string {
	return [][]string{
		[]string{"ZhName"},
	}
}

func ComicCateList(pageSize, page int) ([]*ComicCategoryModel, int64) {
	offset := (page - 1) * pageSize

	query := orm.NewOrm().QueryTable(TbNameComicCategory())
	total, _ := query.Count()

	data := make([]*ComicCategoryModel, 0)
	query.Limit(pageSize, offset).All(&data)

	return data, total
}

func GetComicCateByName(comicCateName string) ComicCategoryModel {
	user := ComicCategoryModel{ZhName: comicCateName}
	o := orm.NewOrm()
	o.Read(&user, "ZhName")
	return user
}

func UpdateComicCate(user *ComicCategoryModel, cols string) error {
	o := orm.NewOrm()
	_, err := o.Update(user, cols)
	return err
}
