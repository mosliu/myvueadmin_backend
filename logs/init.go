package logs

import (
    "github.com/sirupsen/logrus"
    "github.com/shiena/ansicolor"
    "os"
    "github.com/astaxie/beego"
)

// 你可以创建很多instance
//Log to stdout.
var Log = logrus.New()

// Log to File.
var LogF = logrus.New()

func init() {
    logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true,})
    logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
    levelConsole := beego.AppConfig.DefaultString("levelConsole", "debug")
    levelFile := beego.AppConfig.DefaultString("levelFile", "info")
    initLogger(levelConsole, levelFile)
}

func initLogger(levelConsole string, levelFile string) {

    lvlConsole, err := logrus.ParseLevel(levelConsole)
    if err != nil {
        Log.Fatal(err)
        Log.SetLevel(logrus.DebugLevel)
    } else {
        Log.SetLevel(lvlConsole)
    }

    // force colors on for TextFormatter
    Log.Formatter = &logrus.TextFormatter{ForceColors: true,}
    // then wrap the Log output with it
    // 用于解决windows的terminal中彩色不正确的问题
    Log.Out = ansicolor.NewAnsiColorWriter(os.Stdout)

    //init LogF
    fileLocation := beego.AppConfig.DefaultString("FileLocation", "./logrus.Log")
    file, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        LogF.Out = file
    } else {
        Log.Warn("Failed to Log to file, using default stderr")
    }

    lvlFile, err := logrus.ParseLevel(levelFile)
    if err != nil {
        Log.Fatal(err)
        Log.SetLevel(logrus.InfoLevel)
    } else {
        Log.SetLevel(lvlFile)
    }
}

func Assemble() {
    logrus.AddHook(NewContextHook())
    Log.AddHook(NewContextHook())
    LogF.AddHook(NewContextHook())
}
