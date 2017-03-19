package fcontrollers

import (
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Breadcrumb     map[string]string
	BreadcrumbSort []string
}
type menu struct {
	Label     string
	IconClass string
	Active    string
	Url       string
	Items     []menu
}

// Prepare runs after Init before request function execution.
func (this *BaseController) Prepare() {
	this.Layout = "home/layout/layout.html"
	this.Data["D"] = time.Now()
	this.Data["Menus"] = this.getMenuItems()
	this.Data["PageTitle"] = "Tsun Blog"
}

// Finish runs after request function execution.
func (this *BaseController) Finish() {
}

func (this *BaseController) getMenuItems() []menu {
	menus := make([]menu, 0)

	menus = append(menus, menu{
		Label:     "Dashboard",
		IconClass: "fa fa-dashboard",
		Active:    "dashboard",
		Url:       "/",
		Items:     nil})

	// categories := models.GetAllCategories()
	// catetegoryMenus := make([]menu, 0)
	// for _, category := range categories {
	// 	idStr := strconv.Itoa(int(category.ID))
	// 	categoryMenu := menu{
	// 		Label:     category.Name,
	// 		IconClass: "fa fa-circle-o",
	// 		Active:    category.Name,
	// 		Url:       "/blog?category=" + idStr,
	// 		Items:     nil}
	// 	catetegoryMenus = append(catetegoryMenus, categoryMenu)
	// }
	// categoryMenu := menu{
	// 	Label:     "Category",
	// 	IconClass: "fa fa-bus",
	// 	Active:    "category",
	// 	Url:       "#",
	// 	Items:     catetegoryMenus}

	// menus = append(menus, categoryMenu)

	return menus
}
