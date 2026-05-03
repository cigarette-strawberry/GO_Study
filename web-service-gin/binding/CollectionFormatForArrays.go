package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	你可以使用表单绑定中的 collection_format 结构体标签来控制 Gin 如何拆分切片/数组字段的列表值。

	multi（默认）：分号分隔的值
	csv：分号分隔的值
	ssv：空格分隔的值
	tsv：制表符分隔的值
	pipes：管道符分隔的值

	在 form 标签中使用 default 来设置回退值。
	对于 multi 和 csv，使用分号分隔默认值：default=1;2;3。
	对于 ssv、tsv 和 pipes，在默认值中使用自然分隔符。
*/

type Filters struct {
	Tags   []string `form:"tags,default=go;web;api" collection_format:"csv"`         // /search?tags=go,web,api
	Labels []string `form:"labels,default=bug;helpwanted" collection_format:"multi"` // /search?labels=bug&labels=helpwanted
	IdsSSV []int    `form:"ids_ssv,default=1 2 3" collection_format:"ssv"`           // /search?ids_ssv=1 2 3
	IdsTSV []int    `form:"ids_tsv,default=1\t2\t3" collection_format:"tsv"`         // /search?ids_tsv=1\t2\t3
	Levels []int    `form:"levels,default=1|2|3" collection_format:"pipes"`          // /search?levels=1|2|3
}

func CollectionFormatForArrays(c *gin.Context) {
	var filters Filters
	if err := c.ShouldBind(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, filters)
}
