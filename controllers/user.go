package controllers

import (
	"NativeAppApi/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	beego.Controller
}

type aaa struct {
	Name string
}

func (this *UserController) CheckUsenameExists() {
	userName := this.Ctx.Input.Param(":userName")
	o := orm.NewOrm()
	user := models.User{UserName: userName}

	if success := o.Read(&user, "UserName"); success == nil {
		return
	}

	this.Data["json"] = true
	this.ServeJSON()
}

func (this *UserController) CreateNewUser() {
	userName := this.GetString("userName")
	phoneNum := this.GetString("phoneNum")
	password := this.GetString("password")
	birthday := this.GetString("birthday")

	//パスワードハッシュ化する
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	hashPassword := string(hash)

	//必ず山岳部標準時にする必要がある
	layout := "2006-01-02"
	parseBarthday, _ := time.Parse(layout, birthday)

	jwtToken := IssueJWT(userName, phoneNum, parseBarthday)

	user := models.User{}
	user.UserName = userName
	user.PhoneNumber = phoneNum
	user.Password = hashPassword
	user.Birthday = parseBarthday
	user.JwtToken = jwtToken

	o := orm.NewOrm()
	_, err := o.Insert(&user)
	if err != nil {
		return
	}

	this.Data["json"] = jwtToken
	this.ServeJSON()
}

func (this *UserController) Login() {
	phoneNum := this.GetString("phoneNum")
	password := this.GetString("password")

	user := models.User{PhoneNumber: phoneNum}
	o := orm.NewOrm()
	if err := o.Read(&user, "PhoneNumber"); err != nil {
		return
	}

	passwordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passwordError != nil {
		return
	}

	jwtToken := IssueJWT(user.UserName, user.PhoneNumber, user.Birthday)
	user.JwtToken = jwtToken
	o.Update(&user, "JwtToken")
	this.Data["json"] = jwtToken
	this.ServeJSON()
}
