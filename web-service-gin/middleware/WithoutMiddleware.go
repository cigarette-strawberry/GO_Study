package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Gin 提供了两种创建路由引擎的方式，区别在于默认附加了哪些中间件。

	gin.Default() — 带有 Logger 和 Recovery 创建一个已附加两个中间件的路由器
		Logger — 将请求日志写入标准输出（方法、路径、状态码、延迟）
		Recovery — 从处理函数中的任何 panic 恢复并返回 500 响应，防止服务器崩溃
	gin.New() — 一个空白引擎 创建一个完全空白的路由器，不附加任何中间件
		你想使用结构化日志记录器（如 slog 或 zerolog）代替默认的文本日志记录器
		你想自定义 panic 恢复行为
		你正在构建一个需要最小化或专用中间件栈的微服务
*/
func WithoutMiddleware(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
