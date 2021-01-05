package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type aaa struct {
	Name string
}

func (this *UserController) Test() {
	log.Println("Hello")
	a := aaa{}
	a.Name = "じゅんや"

	this.Data["json"] = a
	this.ServeJSON()
}
