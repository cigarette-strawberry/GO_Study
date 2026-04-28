package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	将默认值放在表单键之后：form:"name,default=cigarette"。
	对于集合，使用 collection_format:"multi|csv|ssv|tsv|pipes" 指定如何拆分值。
	对于 multi 和 csv，在默认值中使用分号分隔值（例如 default=1;2;3）。Gin 会在内部将其转换为逗号，以保持标签解析器的明确性。
	对于 ssv（空格）、tsv（制表符）和 pipes（|），在默认值中使用自然分隔符。
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
