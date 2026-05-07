package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	你可以使用 LoggerConfig 跳过特定路径或基于自定义逻辑的日志记录
		SkipPaths 排除特定路由的日志记录——适用于会产生噪音的健康检查或指标端点
		Skip 是一个接收 *gin.Context 并返回 true 来跳过日志记录的函数——适用于条件逻辑，如跳过成功响应的日志
*/

func SkipLogging(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "SkipLogging",
	})
}
