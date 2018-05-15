package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/mosliu/myvueadmin_backend/controllers:UserController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/user/info`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
