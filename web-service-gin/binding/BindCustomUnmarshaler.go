package main

import (
	"encoding"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
要覆盖 Gin 的默认绑定逻辑，请在你的类型上定义一个满足 Go 标准库中 encoding.TextUnmarshaler 接口的函数。然后在要绑定的字段的 uri/form 标签中指定 parser=encoding.TextUnmarshaler
*/
type Birthday string

func (b *Birthday) UnmarshalText(text []byte) error {
	*b = Birthday(strings.Replace(string(text), "-", "/", -1))
	return nil
}

var _ encoding.TextUnmarshaler = (*Birthday)(nil)

var request struct {
	Birthday         Birthday   `form:"birthday,parser=encoding.TextUnmarshaler"`
	Birthdays        []Birthday `form:"birthdays,parser=encoding.TextUnmarshaler" collection_format:"csv"`
	BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02,parser=encoding.TextUnmarshaler" collection_format:"csv"`
}

func BindCustomUnmarshaler(c *gin.Context) {
	_ = c.BindQuery(&request)
	c.JSON(http.StatusOK, request)
}
