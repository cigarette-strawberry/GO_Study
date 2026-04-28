package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
