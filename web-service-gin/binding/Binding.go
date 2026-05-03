package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Gin 提供了强大的绑定系统，可以将请求数据解析到 Go 结构体中并自动验证。
你无需手动调用 c.PostForm() 或读取 c.Request.Body，只需定义一个带标签的结构体，让 Gin 来完成工作

# Gin 提供了两组绑定方法 Bind 与 ShouldBind

大多数情况下，推荐使用 ShouldBind 以获得更好的错误处理控制

（json、xml、yaml、form、uri、header）来映射字段。验证规则放在 binding 标签中
*/
type Form struct {
	User     string `form:"user" bind:"required"`
	Password string `form:"password" bind:"required"`
}

func Binding(c *gin.Context) {
	var form Form
	// ShouldBind checks Content-Type to select a binding engine automatically
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":     form.User,
		"password": form.Password,
	})
}
