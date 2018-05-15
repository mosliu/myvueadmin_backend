package controllers

import "fmt"

type MainController struct {
	//beego.Controller
    BaseController
}

func (c *MainController) Get() {
    fmt.Println("do get method")
	c.Data["Website"] = "beego.me"
    c.Data["Email"] = "astaxie@gmail.com"

	//c.TplName = "1.html"
	c.TplName = "index.tpl"
	//c.display("1")
}
