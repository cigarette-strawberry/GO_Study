package main

import "github.com/gin-gonic/gin"

/*
	使用安全头来保护你的 Web 应用免受常见安全漏洞非常重要。此示例展示了如何向 Gin 应用添加安全头，以及如何避免 Host Header 注入相关攻击（SSRF、开放重定向）

	gin-helmet 是一个为 Go 语言 Web 框架提供 HTTP 安全中间件的集合。 它的核心目标是通过设置各种安全的 HTTP 响应头，来帮助你的 Web 应用防范常见的 Web 安全漏洞。
	你可以把它看作是 Node.js/Express 生态中著名的安全中间件 Helmet 在 Go 语言世界的移植和扩展版本。
*/

func SecurityHeaders(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
