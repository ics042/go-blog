package bcontrollers

import "github.com/ics042/go-blog/models"

type InitDbController struct {
	BaseController
}

func (this *InitDbController) Get() {
	err := models.SyncDb()
	this.Data["ModuleName"] = "system"
	this.Data["MenuName"] = "sync_db"
	if err == nil {
		this.Data["Result"] = "Congratulations, DB synchronization has been done successfully."
	} else {
		this.Data["Result"] = "Sorry, Something wrong during DB synchronization."
	}
	this.TplName = "admin/system/init_db.html"
}
