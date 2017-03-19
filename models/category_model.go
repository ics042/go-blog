package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	ID           int64  `orm:"column(id);pk;auto"`
	Name         string `orm:"column(name)"`
	DisplayOrder int
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt    time.Time `orm:"auto_now;type(datetime)"`
	Blogs        []*Blog   `orm:"reverse(many)" json:"-"`
}

func (this *Category) TableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func GetAllCategories() (categories []*Category) {
	o := orm.NewOrm()
	o.QueryTable("category").All(&categories)
	return categories
}

func GetCategory(id int64) (category Category) {
	o := orm.NewOrm()
	category = Category{ID: id}
	o.Read(&category)
	return category
}

func AddCategory(name, displayOrder string) error {
	o := orm.NewOrm()
	displayOrderI, err := strconv.Atoi(displayOrder)

	category := &Category{
		Name:         name,
		DisplayOrder: displayOrderI}
	_, err = o.Insert(category)

	return err
}

func EditCategory(id int64, name string, displayOrder int) error {
	o := orm.NewOrm()
	category := Category{ID: id}
	var err error
	if o.Read(&category) == nil {
		category.Name = name
		category.DisplayOrder = displayOrder
		_, err = o.Update(&category)
	}
	return err
}

func DeleteCategory(id int64) error {
	o := orm.NewOrm()
	category := Category{ID: id}
	_, err := o.Delete(&category)
	return err
}
