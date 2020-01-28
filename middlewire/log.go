package middlewire

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gotest/config"
	"os"
	"path"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.Log_FILE_PATH
	logFileName := config.Log_FILE_NAME

	fileName := path.Join(logFilePath, logFileName)

	src, err := os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	logger := logrus.New()
	logger.Out = src
 	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})

	return func(c *gin.Context) {

		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		statusCode := c.Writer.Status()

		ClientIp := c.ClientIP()

		logger.WithFields(logrus.Fields{"status_code": statusCode, "latency_time" : latencyTime,
			"req_method"  : reqMethod,
			"req_uri" : reqUri,
			"client_ip" : ClientIp}).Info()

	}

}



