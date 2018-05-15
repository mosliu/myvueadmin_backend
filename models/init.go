package models

import (
    "github.com/astaxie/beego"
    "net/url"
    "github.com/astaxie/beego/orm"

    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "github.com/mosliu/myvueadmin_backend/logs"
    "github.com/sirupsen/logrus"
)

var log = logs.Log.WithFields(logrus.Fields{
    "pkg":"models",
})
func Connect() {

    var dsn string

    dbhost := beego.AppConfig.String("db.host")
    dbport := beego.AppConfig.String("db.port")
    dbuser := beego.AppConfig.String("db.user")
    dbpassword := beego.AppConfig.String("db.password")
    dbname := beego.AppConfig.String("db.name")
    timezone := beego.AppConfig.String("db.timezone")
    if dbport == "" {
        dbport = "3306"
    }
    orm.RegisterDriver("mysql", orm.DRMySQL)
    dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbuser, dbpassword, dbhost, dbport, dbname)
    // fmt.Println(dsn)

    if timezone != "" {
        dsn = dsn + "&loc=" + url.QueryEscape(timezone)
    }
    orm.RegisterDataBase("default", "mysql", dsn)
    //orm.RegisterModel(
    //     new(Role), new(RoleAuth), new(User),new(Auth),
    //)
    if beego.AppConfig.String("runmode") == "dev" {
        orm.Debug = true
    }
}
