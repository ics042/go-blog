package routers

import (
	"github.com/astaxie/beego"
	"github.com/ics042/go-blog/controllers/bcontrollers"
	"github.com/ics042/go-blog/controllers/fcontrollers"
)

func init() {

	// Backend Routes
	// Login
	beego.Router("/login", &bcontrollers.LoginController{})
	// Init Database
	beego.Router("/admin/init-db", &bcontrollers.InitDbController{})
	// Admin
	beego.Router("/admin", &bcontrollers.AdminController{})
	// Category
	beego.Router("/category", &bcontrollers.CategoryController{})
	// Blog
	beego.Router("/blog", &bcontrollers.BlogController{})

	// Fronend Routes
	// Home Page
	beego.Router("/", &fcontrollers.HomeController{})
	beego.Router("/home", &fcontrollers.HomeController{})
}
