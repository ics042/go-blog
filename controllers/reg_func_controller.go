package controllers

import (
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
)

func GetStatus(in int) (out string) {

	switch in {
	case 1:
		out = "Active"
	default:
		out = "Inactive"
	}
	return out
}
func GetBlogStatus(in int8) (out string) {

	switch in {
	case 1:
		out = "Released"
	default:
		out = "Draft"
	}
	return out
}

func Markdown(input string) (out string) {

	return string(blackfriday.MarkdownCommon([]byte(input)))
}

func GetHundred(input string) (out string) {

	if len(input) > 200 {
		return input[0:200]
	} else {
		return input
	}

}

func RegisterTemplateFuncs() {
	beego.AddFuncMap("GetStatus", GetStatus)
	beego.AddFuncMap("GetBlogStatus", GetBlogStatus)
	beego.AddFuncMap("Markdown", Markdown)
	beego.AddFuncMap("GetHundred", GetHundred)
}
