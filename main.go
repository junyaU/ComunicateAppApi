package main

import (
	"NativeAppApi/models"
	_ "NativeAppApi/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	setupDB()

	sessionconf := &session.ManagerConfig{
		CookieName:      "NativeID",
		Gclifetime:      3600,
		EnableSetCookie: true,
		Maxlifetime:     3600,
		Secure:          true,
		CookieLifeTime:  3600,
		ProviderConfig:  "",
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.Run()
}

func setupDB() {
	db := beego.AppConfig.String("driver")

	orm.RegisterDriver(db, orm.DRMySQL)
	orm.RegisterDataBase("default", db, beego.AppConfig.String("sqlconn")+"?charset=utf8")
	orm.RegisterModel(
		new(models.User),
	)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}
