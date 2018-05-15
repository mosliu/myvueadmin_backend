package main

import (
	"github.com/beego/bee/logger/colors"
	_ "github.com/mosliu/myvueadmin_backend/routers"

	"os"
	"github.com/astaxie/beego"

	"strings"

	"github.com/mosliu/myvueadmin_backend/filters"
	"github.com/mosliu/myvueadmin_backend/logs"
	"github.com/mosliu/myvueadmin_backend/models"
	"github.com/sirupsen/logrus"
    "github.com/mosliu/myvueadmin_backend/libs"
)

var log *logrus.Entry
func init() {
    //装配logs配置
    logs.Assemble()
    log = logs.Log.WithFields(logrus.Fields{})
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//log.Println(GetAppPath())
	//log.Println(os.Getwd())

	if pwd, err := os.Getwd(); err == nil {
		apppath := libs.GetAppPath()
		if strings.Compare(pwd, apppath) == 0 {
			log.Info("运行目录与实际目录一致，不执行操作")
		} else {
			log.Info(colors.Green("运行目录与实际目录不一致，进行设定"))
			//  \src\github.com\mosliu\myvueadmin_backend\conf\app.conf
			log.Info(pwd + "\\src\\github.com\\mosliu\\myvueadmin_backend\\conf\\app.conf")
			beego.LoadAppConfig("ini", pwd+"\\src\\github.com\\mosliu\\myvueadmin_backend\\conf\\app.conf")
			beego.SetViewsPath(pwd + "\\src\\github.com\\mosliu\\myvueadmin_backend\\views")
		}
	}


	// 装配filters
	filters.Assemble()
	// 连接数据库
	models.Connect()

	beego.Run()
}

