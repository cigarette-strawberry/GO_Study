package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	Gin 内置了 gin.BasicAuth() 中间件，实现了 HTTP 基本认证。它接受一个 gin.Accounts 映射（map[string]string 的快捷方式），包含用户名/密码对，并保护应用它的任何路由组

	c.MustGet(key string) interface{} 接受一个字符串类型的 key，并尝试从请求上下文 c 中获取与该 key 关联的值。
	如果该值存在，c.MustGet 会返回这个值，类型为 interface{}，需要进行类型断言才能转换为具体类型；如果值不存在，程序会发生 panic

	c.Get(key string) (interface{}, bool) 也用于从请求上下文获取值，但它不会 panic。
	如果值存在，它返回值和 true；如果值不存在，它返回 nil 和 false

	c.MustGet 适用于你确定某个值在上下文中一定存在的场景，使用它可以简化代码，避免繁琐的存在性检查；
	而 c.Get 则更适合值可能不存在的情况，通过返回的布尔值可以进行更灵活的处理
*/

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func UsingBasicauthMiddleware(c *gin.Context) {
	// 获取用户信息，该信息是由基本身份验证中间件设置的
	// 使用 MustGet 获取值
	user := c.MustGet(gin.AuthUserKey).(string)
	// 类型断言将 interface{} 转换为具体类型
	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}
