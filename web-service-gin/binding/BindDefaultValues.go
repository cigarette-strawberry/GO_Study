package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	有时你希望当客户端没有发送值时，字段能回退到默认值。Gin 的表单绑定通过 form 结构体标签中的 default 选项支持默认值。这适用于标量值，从 Gin v1.11 开始，也适用于具有显式集合格式的集合（切片/数组）。

	将默认值放在表单键之后：form:"name,default=cigarette"。
	对于集合，使用 collection_format:"multi|csv|ssv|tsv|pipes" 指定如何拆分值。
	对于 multi 和 csv，在默认值中使用分号分隔值（例如 default=1;2;3）。Gin 会在内部将其转换为逗号，以保持标签解析器的明确性。
	对于 ssv（空格）、tsv（制表符）和 pipes（|），在默认值中使用自然分隔符。

	对于 multi 和 csv，分号用于分隔默认值；不要在这些格式的单个默认值中包含分号。
	无效的 collection_format 值将导致绑定错误
*/

type Info struct {
	Name      string    `form:"name,default=cigarette"`
	Age       int       `form:"age,default=10"`
	Friends   []string  `form:"friends,default=Will;Bill"` // multi/csv: use ; in defaults
	Addresses [2]string `form:"addresses,default=foo bar" collection_format:"ssv"`
	LapTimes  []int     `form:"lap_times,default=1;2;3" collection_format:"csv"`
}

func BindDefaultValues(c *gin.Context) {
	var info Info
	if err := c.ShouldBind(&info); err != nil { // infers binder by Content-Type
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, info)
}
