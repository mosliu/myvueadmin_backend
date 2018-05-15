package controllers

import (
    "strings"
    "github.com/astaxie/beego"
)

type BaseController struct {
    beego.Controller
    controllerName string
    actionName     string

    userId         int
}

var form map[string]interface{}

func (this *BaseController) Prepare() {
    controllerName, actionName := this.GetControllerAndAction()
    this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
    this.actionName = strings.ToLower(actionName)
    log.Debug("controllerName:" + this.controllerName)

    this.Data["version"] = beego.AppConfig.String("version")
    this.Data["siteName"] = beego.AppConfig.String("site.name")

    this.Data["curRoute"] = this.controllerName + "." + this.actionName
    this.Data["curController"] = this.controllerName
    this.Data["curAction"] = this.actionName
    this.doAuthCheck()
}
//登录权限验证
func (this *BaseController) doAuthCheck() {
    //arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
    //this.userId = 0
    //if len(arr) == 2 {
    //    idstr, password := arr[0], arr[1]
    //    userId, _ := strconv.Atoi(idstr)
    //    if userId > 0 {
    //        user, err := models.AdminGetById(userId)
    //        if err == nil && password == libs.Md5([]byte(this.getClientIp()+"|"+user.Password+user.Salt)) {
    //            this.userId = user.Id
    //            this.loginName = user.LoginName
    //            this.userName = user.RealName
    //            this.user = user
    //            this.AdminAuth()
    //        }
    //
    //        isHasAuth := strings.Contains(this.allowUrl, this.controllerName+"/"+this.actionName)
    //        noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start/show/ajaxapisave/index/group/public/env/code/apidetail"
    //        isNoAuth := strings.Contains(noAuth, this.actionName)
    //        if isHasAuth == false && isNoAuth == false {
    //            this.Ctx.WriteString("没有权限")
    //            this.ajaxMsg("没有权限", MSG_ERR)
    //            return
    //        }
    //    }
    //}
    //
    //if this.userId == 0 && (this.controllerName != "login" && this.actionName != "loginin") {
    //    this.redirect(beego.URLFor("LoginController.LoginIn"))
    //}
}

func (this *BaseController) Rsp(code int, message string,data ...interface{}) {

    jsonData := make(map[string]interface{}, 3)

    jsonData["code"] = code
    jsonData["message"] = message

    if len(data) > 0 && data[0] != nil {
        jsonData["data"] = data[0]
    }

    //returnJSON, err := json.Marshal(jsonData)
    //if err != nil {
    //    beego.Error(err)
    //}
    //this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
    //this.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
    //io.WriteString(this.Ctx.ResponseWriter, string(returnJSON))
    this.Data["json"] = &jsonData
    this.ServeJSON()
}

//加载模板
func (this *BaseController) display(tpl ...string) {
    log.Debug("this.TplName:" + this.TplName)
    var tplname string
    if len(tpl) > 0 {
        log.Debug("tpl:" + tpl[0])
        tplname = strings.Join([]string{tpl[0], "html"}, ".")
    } else {
        tplname = this.controllerName + "/" + this.actionName + ".html"
    }
    //this.Layout = "public/layout.html"
    log.Debug("tplname:" + tplname)
    this.TplName = tplname
}

// 是否POST提交
func (self *BaseController) isPost() bool {
    return self.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
    s := self.Ctx.Request.RemoteAddr
    l := strings.LastIndex(s, ":")
    return s[0:l]
}
