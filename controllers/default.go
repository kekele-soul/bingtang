package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"kekele/models"
)

type MainController struct {
	beego.Controller
}
//Get请求
func (c *MainController) Get() {
	//获取数据
 	user := c.Ctx.Input.Query("feiyue")
 	psd := c.Ctx.Input.Query("0123")
	if user != "xiaoliu" && psd != "12345" {
		c.Ctx.ResponseWriter.Write([]byte("请求错误，请重新尝试"))
		return
	}
 		c.Ctx.ResponseWriter.Write([]byte("欢迎来到飞跃的页面"))
}
//func (c *MainController) Post() {
//	name := c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	if name != "xiaomeng" && age != "20" {
//		c.Ctx.WriteString("请求错误，请重新尝试。")
//		return
//	}
//	c.Ctx.WriteString("欢迎进入小梦的页面")
//}
func (c *MainController) Post(){
	//解析json格式
	var person models.Commty
//读取前端页面
	dataBytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("接受错误")
		return
	}
	//解析前端页面
	err =json.Unmarshal(dataBytes,&person)
	if err != nil {
		c.Ctx.WriteString("解析错误")
		return
	}
	fmt.Println("姓名",person.Name)
	fmt.Println("年龄",person.Age)
	fmt.Println("性别",person.Sex)
}