package bcontrollers

type AdminController struct {
	BaseController
}

func (this *AdminController) Get() {
	this.TplName = "admin/home.html"
	this.LayoutSections = make(map[string]string)
	// this.LayoutSections["HtmlHead"] = "blogs/html_head.tpl"
	// this.LayoutSections["Scripts"] = "blogs/scripts.tpl"
	this.Data["PageTitle"] = "Home"
	this.Breadcrumb = make(map[string]string)
	this.Breadcrumb["Home"] = "/admin"
	this.Data["Breadcrumb"] = this.Breadcrumb
	// this.Data["ModuleName"] = "product"
	// this.Data["MenuName"] = "log_out"

}
