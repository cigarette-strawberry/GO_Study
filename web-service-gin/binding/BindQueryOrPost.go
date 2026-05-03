package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
ShouldBind 会根据 HTTP 方法和 Content-Type 请求头自动选择绑定引擎：

对于 GET 请求，使用查询字符串绑定（form 标签）。
对于 POST/PUT 请求，它会检查 Content-Type——对 application/json 使用 JSON 绑定，对 application/xml 使用 XML 绑定，对 application/x-www-form-urlencoded 或 multipart/form-data 使用表单绑定。

这意味着单个处理函数可以同时接受来自查询字符串和请求体的数据，无需手动选择数据源

格式 2006-01-02 表示”年-月-日”。time_utc:"1" 标签确保解析后的时间为 UTC 时区
*/
type PersonInfo struct {
	Name     string `form:"name"`
	Address  string `form:"address"`
	Birthday string `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func BindQueryOrPost(c *gin.Context) {
	var person_info PersonInfo

	if err := c.ShouldBind(&person_info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Name: %s, Address: %s, Birthday: %s\n", person_info.Name, person_info.Address, person_info.Birthday)
	c.JSON(http.StatusOK, gin.H{
		"name":     person_info.Name,
		"address":  person_info.Address,
		"birthday": person_info.Birthday,
	})
}
