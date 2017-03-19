package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func SyncDb() error {
	err := orm.RunSyncdb("default", false, true)
	return err
}

func InitDb() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/golang_blog?charset=utf8&loc=Pacific%2FAuckland")
	orm.RegisterModel(new(Blog))
	orm.RegisterModel(new(Category))
}
