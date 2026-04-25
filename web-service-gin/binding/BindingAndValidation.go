package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	Gin 提供了两组绑定方法
	类型 - Must bind
		方法 - Bind、BindJSON、BindXML、BindQuery、BindYAML
		行为 - 这些方法底层使用 MustBindWith。如果存在绑定错误，请求将使用 c.AbortWithError(400, err).SetType(ErrorTypeBind) 中止。这会将响应状态码设置为 400，并将 Content-Type 头设置为 text/plain; charset=utf-8。注意，如果你在此之后尝试设置响应码，将会出现警告 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422。如果你希望更好地控制行为，请考虑使用 ShouldBind 等效方法。
	类型 - Should bind
		方法 - ShouldBind、ShouldBindJSON、ShouldBindXML、ShouldBindQuery、ShouldBindYAML
		行为 - 这些方法底层使用 ShouldBindWith。如果存在绑定错误，错误会被返回，由开发者负责适当地处理请求和错误。

	使用 Bind 方法时，Gin 会尝试根据 Content-Type 头来推断绑定器。如果你确定要绑定的内容类型，可以使用 MustBindWith 或 ShouldBindWith。
*/

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginJSON(c *gin.Context) {
	var json Login

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "menu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func LoginXML(c *gin.Context) {
	var xml Login

	if err := c.ShouldBindXML(&xml); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if xml.User != "menu" || xml.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func LoginForm(c *gin.Context) {
	var form Login

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.User != "menu" || form.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
