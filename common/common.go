package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gotest/config"
	"net/http"
	"os"
	"path"
)
// 返回JSON
func RetJson(code, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg"  : msg,
		"data" : data,
	})
	c.Abort()
}
var log = logrus.New()
func  Logger()  *logrus.Logger {
	logFilePath := config.Log_FILE_PATH
	logFileName := config.Log_FILE_NAME
	fileName := path.Join(logFilePath, logFileName)

	src, err := os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	log.Out = src
	log.Formatter = &logrus.JSONFormatter{}
	return log
}