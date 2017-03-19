package fcontrollers

import (
	"strconv"

	"github.com/ics042/go-blog/models"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	// this.Data["PageTitle"] = "Dashboard"
	// this.Breadcrumb = make(map[string]string)
	// this.Breadcrumb["Dashboard"] = "/"
	// this.BreadcrumbSort = []string{"Dashboard"}
	this.Data["ModuleName"] = "dashboard"
	this.Data["MenuName"] = "dashboard"

	action := this.Input().Get("action")
	switch action {
	case "detail":
		this.Detail()
	default:
		this.Index()
	}
	this.Data["Breadcrumb"] = this.Breadcrumb
	this.Data["BreadcrumbSort"] = this.BreadcrumbSort
}
func (this *HomeController) Index() {
	this.TplName = "home/dashboard.html"
	blogKeys, blogMap := models.GetDashBoardBlogs()
	this.Data["Blogs"] = blogMap
	this.Data["BlogKeys"] = blogKeys
}

func (this *HomeController) Detail() {
	this.TplName = "home/detail.html"
	idStr := this.Input().Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	blog := models.GetBlog(id)
	this.Data["Blog"] = blog
}
