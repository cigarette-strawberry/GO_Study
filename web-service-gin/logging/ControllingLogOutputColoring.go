package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	控制日志输出着色
		默认情况下，控制台上的日志输出应根据检测到的 TTY 进行着色

		gin.DisableConsoleColor() // 禁用日志着色
		gin.ForceConsoleColor() // 启用日志着色

		日志上色 不能使用 自定日志
*/

func ControllingLogOutputColoring(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "ControllingLogOutputColoring",
	})
}
