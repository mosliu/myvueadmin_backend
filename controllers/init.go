package controllers
import (
    "github.com/mosliu/myvueadmin_backend/logs"
    "github.com/sirupsen/logrus"
)

var log = logs.Log.WithFields(logrus.Fields{
    "pkg":"controllers",
})