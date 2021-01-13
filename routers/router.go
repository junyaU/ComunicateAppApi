// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"NativeAppApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/sendSms", &controllers.SmsController{}, "post:SendMessage")
	beego.Router("/userNameCheck/:userName:string", &controllers.UserController{}, "get:CheckUsenameExists")
	beego.Router("/createUser", &controllers.UserController{}, "post:CreateNewUser")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
}
