package bcontrollers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/ics042/go-blog/models"
)

type BlogController struct {
	BaseController
}

func (this *BlogController) Post() {
	tid := this.Input().Get("cId")
	title := this.Input().Get("title")
	categoryIdStr := this.Input().Get("categoryId")
	content := this.Input().Get("content")

	var err error
	if tid == "" {
		categoryId, _ := strconv.ParseInt(categoryIdStr, 10, 64)
		err = models.AddBlog(title, categoryId, content)
	} else {
		id, _ := strconv.ParseInt(tid, 10, 64)
		categoryId, _ := strconv.ParseInt(categoryIdStr, 10, 64)
		err = models.EditBlog(id, title, categoryId, content)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/blog", 302)
}

func (this *BlogController) Get() {
	this.Data["PageTitle"] = "Blog"
	this.Breadcrumb = make(map[string]string)
	this.Breadcrumb["Home"] = "/admin"
	this.Breadcrumb["Blog"] = "/blog"
	this.BreadcrumbSort = []string{"Home", "Blog"}
	this.Data["ModuleName"] = "blog"
	this.Data["MenuName"] = "manage_blog"

	var tempMap map[string][]string
	tempMap = this.Input()
	if tempMap["action"] == nil {

	}
	action := this.GetString("action")
	switch action {
	case "create":
		this.Create()
	case "edit":
		this.Edit()
	case "del":
		this.Delete()
	case "updateStatus":
		this.UpdateStatus()
	default:
		this.Index()
	}
	this.Data["Breadcrumb"] = this.Breadcrumb
	this.Data["BreadcrumbSort"] = this.BreadcrumbSort
}
func (this *BlogController) Create() {
	this.TplName = "admin/blog/form.html"
	this.Data["Categories"] = models.GetAllCategories()
	this.Data["Action"] = "Add"
	this.Breadcrumb["Add Blog"] = "#"
	this.BreadcrumbSort = []string{"Home", "Blog", "Add Blog"}
}
func (this *BlogController) Edit() {
	id, _ := strconv.ParseInt(this.Input().Get("id"), 10, 64)
	this.TplName = "admin/blog/form.html"
	this.Data["Categories"] = models.GetAllCategories()
	this.Data["Blog"] = models.GetBlog(id)
	this.Data["Action"] = "Edit"
	this.Breadcrumb["Edit Catetory"] = "#"
	this.BreadcrumbSort = []string{"Home", "Blog", "Edit Blog"}
}
func (this *BlogController) Delete() {
	id, _ := strconv.ParseInt(this.Input().Get("id"), 10, 64)
	err := models.DeleteBlog(id)
	if err != nil {
		this.Data["Error"] = err
	}
	this.Redirect("/blog?action=index", 302)
}
func (this *BlogController) Index() {
	this.TplName = "admin/blog/index.html"
	this.Data["SubTitle"] = "Blog List"
	this.Data["Blogs"] = models.GetAllBlogs()
}
func (this *BlogController) UpdateStatus() {
	id, _ := strconv.ParseInt(this.Input().Get("id"), 10, 64)
	err := models.UpdateBlogStatus(id)
	if err != nil {
		this.Data["Error"] = err
	}
	this.Redirect("/blog?action=index", 302)
}
