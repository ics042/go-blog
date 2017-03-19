package bcontrollers

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
	loginKey := this.GetSession("loginKey")
	if loginKey != nil && loginKey == beego.AppConfig.String("key") {
		this.Layout = "admin/layout/layout.html"
		this.Data["D"] = time.Now()
		this.Data["Menus"] = this.getMenuItems()
	} else {
		this.Redirect("/login", 302)
	}
}

// Finish runs after request function execution.
func (this *BaseController) Finish() {
}

func (this *BaseController) getMenuItems() []menu {
	menus := make([]menu, 0)

	newPostMenu := menu{
		Label:     "Manage Blog",
		IconClass: "fa fa-plus-square",
		Active:    "manage_blog",
		Url:       "/blog?action=index",
		Items:     nil}
	menus = append(menus, newPostMenu)

	systemMenus := make([]menu, 0)
	systemMenus = append(systemMenus, menu{
		Label:     "Manage Category",
		IconClass: "fa fa-circle-o",
		Active:    "manage_category",
		Url:       "/category?action=index",
		Items:     nil})
	systemMenus = append(systemMenus, menu{
		Label:     "Synchronize Database",
		IconClass: "fa fa-circle-o",
		Active:    "sync_db",
		Url:       "/admin/init-db",
		Items:     nil})
	systemMenu := menu{
		Label:     "System",
		IconClass: "fa fa-address-book",
		Active:    "system",
		Url:       "#",
		Items:     systemMenus}
	exitMenu := menu{
		Label:     "Lout Out",
		IconClass: "fa fa-sign-out",
		Active:    "log_out",
		Url:       "/login?action=logout",
		Items:     nil}
	menus = append(menus, systemMenu)
	menus = append(menus, exitMenu)

	return menus
}
