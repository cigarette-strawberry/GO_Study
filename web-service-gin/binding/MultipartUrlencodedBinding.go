package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	ShouldBind 会自动检测 Content-Type，并将 multipart/form-data 或 application/x-www-form-urlencoded 请求体绑定到结构体中。
	使用 form 结构体标签将表单字段名映射到结构体字段，使用 binding:"required" 来强制必填字段。

	ShouldBind 会自动检测 Content-Type，并将 multipart/form-data 或 application/x-www-form-urlencoded 请求体绑定到结构体中。使用 form 结构体标签将表单字段名映射到结构体字段，使用 binding:"required" 来强制必填字段
*/

type LoginFormItem struct {
	User     string `form:"user" json:"user"`
	Password string `form:"password" json:"password"`
}

func MultipartUrlencodedBinding(c *gin.Context) {
	var form LoginFormItem
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.User == "user" && form.Password == "password" {
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	}
}
