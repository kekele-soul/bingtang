package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"kekele/db_mysql"
	"kekele/models"
)

type RegisterController struct {
	beego.Controller
}


//接受数据
func(r *RegisterController) Post(){
	body,err :=ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接受错误，请检查后重试")
		return
	}
	//解析数据
	var user models.Sql
	err = json.Unmarshal(body,&user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析错误")
		return
	}
	//将解析好的数据，进行保存
	_,err = db_mysql.InsertUser(user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败")
		return
	}
	//fmt.Println(id)
	r.Ctx.WriteString("保存成功")
}
