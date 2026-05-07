package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	// 下面两个是 全局中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	middleware := router.Group("/middleware")

	{
		middleware.GET("/withoutMiddleware", WithoutMiddleware)
	}

	{
		middleware.GET("/usingMiddleware", UsingMiddleware)
	}

	// 下面这个是 分组中间件
	// middleware.Use(Logger())
	{
		// 这个则是 路由级中间件
		middleware.GET("/customMiddleware", Logger(), CustomMiddleware)
	}

	middleware.Use(ErrorHandler())
	{
		middleware.GET("/errorHandlingMiddleware", ErrorHandlingMiddleware)
	}

	// 使用 gin.BasicAuth() 中间件进行分组
	// gin.Accounts 是一个用于表示 map[string]string 的简写形式
	authorized := router.Group("/basicauth", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	{
		authorized.GET("/usingBasicauthMiddleware", UsingBasicauthMiddleware)
	}

	{
		middleware.GET("/goroutinesInsideAMiddleware", GoroutinesInsideAMiddleware)
	}

	router.Run(":8080")
}
