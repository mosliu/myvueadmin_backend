// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
    "github.com/mosliu/myvueadmin_backend/controllers"

    "github.com/astaxie/beego"
)

func init() {
    //webns := beego.NewNamespace("web",
    //    beego.NSNamespace("/index",
    //        beego.NSRouter("/1", &controllers.MainController{}),
    //    ),
    //)
    //beego.AddNamespace(webns)
    //ns := beego.NewNamespace("/v1",
    //    beego.NSNamespace("/object",
    //        beego.NSInclude(
    //            &controllers.ObjectController{},
    //        ),
    //    ),
    //    beego.NSNamespace("/user",
    //        beego.NSInclude(
    //            &controllers.UserController{},
    //        ),
    //    ),
    //)
    //beego.AddNamespace(ns)

    beego.Router("/", &controllers.MainController{})
    beego.Include(&controllers.LoginController{})
    beego.Include(&controllers.UserController{})
    //beego.AutoRouter(&controllers.UserController{})
}
