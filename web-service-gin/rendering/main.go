package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	rendering := router.Group("/rendering")

	{
		rendering.GET("/someJSON", SomeJSON)
		rendering.GET("/someXML", SomeXML)
		rendering.GET("/someYAML", SomeYAML)
		rendering.GET("/someProtoBuf", SomeProtoBuf)
	}

	router.Run(":8080")
}
