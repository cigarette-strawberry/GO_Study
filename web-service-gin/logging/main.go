package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 禁用控制台颜色，当将日志写入文件时，您无需使用控制台颜色。
	gin.DisableConsoleColor()

	// 将日志记录到文件中。
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果您需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	gin.DefaultWriter = &lumberjack.Logger{
		Filename:   "gin.log",
		MaxSize:    100, // 兆字节
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	// r := gin.Default()

	/* ------------------------------------上下分隔------------------------------------ */

	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	logging := r.Group("/logging")

	{
		logging.GET("/writeLog", WriteLog)
	}

	{
		logging.GET("/customLogFormat", CustomLogFormat)
	}

	r.Run(":8080")
}
