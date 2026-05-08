package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
	会话允许你跨多个 HTTP 请求存储用户特定的数据。
	由于 HTTP 是无状态的，会话使用 cookie 或其他机制来识别回访用户并检索其存储的数据。

	最简单的方式是将会话数据存储在加密的 cookie 中
*/

func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user", "cigarette")
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func Profile(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	log.Printf("User data: username - %v", user)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
