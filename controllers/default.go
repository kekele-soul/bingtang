package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
 	user := c.Ctx.Input.Query("feiyue")
 	psd := c.Ctx.Input.Query("0123")
	if user != "xiaoliu" && psd != "12345" {
		c.Ctx.ResponseWriter.Write([]byte("请求错误，请重新尝试"))
		return
	}
 		c.Ctx.ResponseWriter.Write([]byte("欢迎来到飞跃的页面"))
}
func (c *MainController) Post() {
	name := c.Ctx.Request.FormValue("name")
	age := c.Ctx.Request.FormValue("age")
	if name != "xiaomeng" && age != "20" {
		c.Ctx.WriteString("请求错误，请重新尝试")
		return
	}
	c.Ctx.WriteString("欢迎进入小梦的页面")
}
//func (c *MainController) Post(){
//
//
//}