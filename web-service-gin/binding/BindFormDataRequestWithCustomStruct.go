package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	Gin 可以自动将表单数据绑定到嵌套结构体中。当你的数据模型由较小的结构体组成时——无论是嵌入字段、指针字段还是匿名内联结构体——Gin 会遍历结构体层次并将每个 form 标签映射到相应的查询参数或表单字段。

	这对于将复杂表单组织成可重用的子结构非常有用，而不是定义一个包含许多字段的扁平结构体。

	所有三种模式——嵌套结构体、嵌套结构体指针和匿名内联结构体——都使用相同的扁平查询参数进行绑定。Gin 不要求参数名称中有任何前缀或嵌套约定。
*/

/*
StructA 是一个结构体类型的定义。它代表了一种数据结构，包含了特定的字段集合。
*StructA 是指向 StructA 类型结构体的指针类型。指针是一个变量，它存储的是另一个变量的内存地址。
要让这个指针指向一个 StructA 实例，你可以使用 new 关键字或者取地址运算符 &

函数接收的是 StructA 类型的实例，对其修改不会影响原始的 a。而 函数接收的是 *StructA 类型的指针，对其修改会影响原始的 a。
*/
type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}
func GetC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}
func GetD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}
