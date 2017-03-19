package bcontrollers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/ics042/go-blog/models"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Post() {
	tid := this.Input().Get("cId")
	name := this.Input().Get("name")
	displayOrder := this.Input().Get("displayOrder")

	var err error
	if tid == "" {
		err = models.AddCategory(name, displayOrder)
	} else {
		id, _ := strconv.ParseInt(tid, 10, 64)
		displayOrderI, _ := strconv.Atoi(displayOrder)
		err = models.EditCategory(id, name, displayOrderI)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/category", 302)
}

func (this *CategoryController) Get() {
	this.Data["PageTitle"] = "Category"
	this.Breadcrumb = make(map[string]string)
	this.Breadcrumb["Home"] = "/admin"
	this.Breadcrumb["Category"] = "/category"
	this.BreadcrumbSort = []string{"Home", "Category"}
	this.Data["ModuleName"] = "system"
	this.Data["MenuName"] = "manage_category"

	action := this.Input().Get("action")
	switch action {
	case "create":
		this.Create()
	case "edit":
		this.Edit()
	case "del":
		this.Delete()
	default:
		this.Index()
	}
	this.Data["Breadcrumb"] = this.Breadcrumb
	this.Data["BreadcrumbSort"] = this.BreadcrumbSort
}
func (this *CategoryController) Create() {
	this.TplName = "admin/category/form.html"
	this.Data["Categories"] = models.GetAllCategories()
	this.Data["Action"] = "Add"
	this.Breadcrumb["Add Catetory"] = "#"
	this.BreadcrumbSort = []string{"Home", "Category", "Add Catetory"}
}
func (this *CategoryController) Edit() {
	id, _ := strconv.ParseInt(this.Input().Get("id"), 10, 64)
	this.TplName = "admin/category/form.html"
	this.Data["Categories"] = models.GetAllCategories()
	this.Data["Category"] = models.GetCategory(id)
	this.Data["Action"] = "Edit"
	this.Breadcrumb["Edit Catetory"] = "#"
	this.BreadcrumbSort = []string{"Home", "Category", "Edit Category"}
}
func (this *CategoryController) Delete() {
	id, _ := strconv.ParseInt(this.Input().Get("id"), 10, 64)
	err := models.DeleteCategory(id)
	if err != nil {
		this.Data["Error"] = err
	}
	this.Redirect("/category?action=index", 301)
}
func (this *CategoryController) Index() {
	this.TplName = "admin/category/index.html"
	this.Data["SubTitle"] = "Category List"
	this.Data["Categories"] = models.GetAllCategories()
}
