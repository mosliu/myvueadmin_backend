package controllers

import (
    "errors"
    m "github.com/mosliu/myvueadmin_backend/models"
    "strings"
    "time"
    "encoding/json"
    "github.com/mosliu/myvueadmin_backend/libs"
)

// LoginController operations for LoginController
type LoginController struct {
    BaseController
}

type loginForm struct {
    Username string `form:"username"`
    Password string `form:"password"`
}
//var form map[string]interface{}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Logout", c.Logout)
}

// Login ...
// @Title Create
// @Description create LoginController
// @Param	body		body 	models.LoginController	true		"body for LoginController content"
// @Success 201 {object} models.LoginController
// @Failure 403 body is empty
// @router /login [post]
func (c *LoginController) Login() {
    //sess := c.StartSession()
    //log.Warn(sess.SessionID())
    //log.Warn(c.CruSession.SessionID())
    log.Info(c.Ctx.Input.Cookie("beegosessionID"))
    //vue axio提交的 isajax是 false
    //logs.Debug(c.Ctx.Input.IsAjax())

    //logs.Debug(c.Ctx.Input.RequestBody)
    //if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
    //    c.Rsp(MSG_ERR, err.Error())
    //}
    logindata := loginForm{}
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &logindata); err != nil {
        c.Rsp(MSG_ERR, err.Error())
    }

    //username := strings.TrimSpace(c.GetString("username"))
    //password := strings.TrimSpace(c.GetString("password"))
    //username := strings.TrimSpace(form["username"].(string))
    //password := strings.TrimSpace(form["password"].(string))
    log.WithField("form",logindata).Debug("login form content")
    username := strings.TrimSpace(logindata.Username)
    password := strings.TrimSpace(logindata.Password)
    //log.Debug(username,password)
    user, err := CheckLogin(username, password)
    if err == nil {
        //accesslist, _ := GetAccessList(user.Id)
        //c.SetSession("accesslist", accesslist)
        user.LastIp = c.getClientIp()
        user.LastLogin = time.Now().Unix()
        user.Update()
        c.SetSession("user", user)
        log.WithField("user",c.GetSession("user")).Debug()

        token := GenerateToken(user,user.LastIp)
        c.SetSession("token", token)

        data := make(map[string]string)
        data["token"] = token
        c.Rsp(MSG_OK, "登录成功", data)

        return
    } else {
        c.Rsp(MSG_NOLOGIN_NAME_ERR, err.Error())
        return
    }
}

func GenerateToken(user m.User,ip string) string {
    token := libs.Str2Md5(user.LoginName + "|" + ip + "|" + user.Password)
    return token
}
//check login
func CheckLogin(username string, password string) (user m.User, err error) {
    if username == "" || password == "" {
        return user,errors.New("账号或密码错误")
    }
    user = m.GetUserByUsername(username)
    if user.Id == 0 {
        return user, errors.New("用户不存在")
    }
    //if user.Password != libs.Pwdhash(password) {
    if user.Password != password {
        return user, errors.New("密码错误")
    }

    if user.Status==-1 {
        return user, errors.New("账号已禁用")
    }
    return user, nil
}

// Login ...
// @Title Create
// @Description create LoginController
// @Param	body		body 	models.LoginController	true		"body for LoginController content"
// @Success 201 {object} models.LoginController
// @Failure 403 body is empty
// @router /logout [post]
func (c *LoginController) Logout() {

    //vue axio提交的 isajax是 false
    //logs.Debug(c.Ctx.Input.IsAjax())
    c.Rsp(MSG_OK, "登出成功")

    return

}