package filters

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/plugins/cors"
)

func assembleCorsFilter() {

    beego.InsertFilter("*",beego.BeforeRouter,cors.Allow(&cors.Options{
        //AllowOrigins:     []string{"https://127.0.0.1"},
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"POST","GET","OPTIONS","DELETE"},
        //AllowHeaders:     []string{"Origin"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin","Access-Control-Allow-Credentials", "content-type","X-Token"},
        //ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))
}
