package main

import (
	"encoding"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
要覆盖 Gin 的默认绑定逻辑，请在你的类型上定义一个满足 Go 标准库中 encoding.TextUnmarshaler 接口的函数。
然后在要绑定的字段的 uri/form 标签中指定 parser=encoding.TextUnmarshaler
注意： 如果为一个没有实现 encoding.TextUnmarshaler 的类型指定了 parser=encoding.TextUnmarshaler，Gin 会忽略它并使用默认绑定逻辑

如果一个类型已经实现了 encoding.TextUnmarshaler，但你希望自定义 Gin 绑定该类型的方式（例如更改返回的错误消息），可以改为实现专用的 BindUnmarshaler 接口
注意： 如果一个类型同时实现了 encoding.TextUnmarshaler 和 BindUnmarshaler，Gin 默认会使用 BindUnmarshaler，除非你在绑定标签中指定了 parser=encoding.TextUnmarshaler
*/

type Birthday string

// 为 Birthday 类型定义了一个方法 UnmarshalText，实现了 encoding.TextUnmarshaler 接口。
// 这个接口用于将文本数据反序列化为自定义类型。
func (b *Birthday) UnmarshalText(text []byte) error {
	// 在 UnmarshalText 方法中，将传入的字节切片 text 转换为字符串，然后使用 strings.Replace 函数将字符串中的所有 - 替换为 /，再将结果转换回 Birthday 类型并赋值给接收者 b。
	*b = Birthday(strings.Replace(string(text), "-", "/", -1))
	return nil
}

// 这行代码是一个类型断言，用于确保 Birthday 类型确实实现了 encoding.TextUnmarshaler 接口。
// 它实际上不会执行任何操作，但如果 Birthday 类型没有正确实现 UnmarshalText 方法，编译器会报错。
var _ encoding.TextUnmarshaler = (*Birthday)(nil)

var text_unmarshaler struct {
	Birthday         Birthday   `form:"birthday,parser=encoding.TextUnmarshaler"`
	Birthdays        []Birthday `form:"birthdays,parser=encoding.TextUnmarshaler" collection_format:"csv"`
	BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02,parser=encoding.TextUnmarshaler" collection_format:"csv"`
}

func TextUnmarshaler(c *gin.Context) {
	_ = c.BindQuery(&text_unmarshaler)
	c.JSON(http.StatusOK, text_unmarshaler)
}

/* ---------------------------------------------------------------------- */

func (b *Birthday) UnmarshalParam(param string) error {
	*b = Birthday(strings.Replace(param, "-", "/", -1))
	return nil
}

var _ binding.BindUnmarshaler = (*Birthday)(nil)

var bind_unmarshaler struct {
	Birthday         Birthday   `form:"birthday"`
	Birthdays        []Birthday `form:"birthdays" collection_format:"csv"`
	BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02" collection_format:"csv"`
}

func BindUnmarshaler(c *gin.Context) {
	_ = c.BindQuery(&bind_unmarshaler)
	c.JSON(http.StatusOK, bind_unmarshaler)
}

/*
	encoding.TextUnmarshaler 和 binding.BindUnmarshaler

	接口定义：encoding.TextUnmarshaler 是 Go 标准库 encoding 包中的一个接口。它定义了一个方法 UnmarshalText，用于将文本数据反序列化为自定义类型。
	用途：当你有一个自定义类型，并且希望能够将字符串形式的文本数据转换为该自定义类型时，可以实现这个接口。
	例如，如果你有一个自定义的日期类型，它可能有特定的格式要求，通过实现 UnmarshalText 方法，可以将符合该格式的字符串转换为你的自定义日期类型。
	UnmarshalText 方法接收一个字节切片 text，这通常是从外部数据源（如文件、网络请求等）读取的文本数据。
	方法需要将这个字节切片转换为自定义类型，并返回可能发生的错误。如果转换成功，应返回 nil 错误。

	接口定义：binding.BindUnmarshaler 是 Gin 框架中的一个接口。它定义了一个方法 UnmarshalParam，主要用于在请求数据绑定过程中，将从请求参数中获取的字符串转换为自定义类型。
	用途：在 Web 开发中，当处理 HTTP 请求时，Gin 框架需要将请求参数（如查询字符串、表单数据等）转换为 Go 结构体中的字段类型。
	如果结构体字段是自定义类型，就可以通过实现 binding.BindUnmarshaler 接口来定义如何进行这种转换。
	UnmarshalParam 方法接收一个字符串 param，这个字符串是从请求参数中获取的值。
	方法需要将这个字符串转换为自定义类型，并返回可能发生的错误。如果转换成功，应返回 nil 错误。

	总结：encoding.TextUnmarshaler 是标准库提供的通用文本反序列化接口，适用于各种需要将文本转换为自定义类型的场景；
	而 binding.BindUnmarshaler 是 Gin 框架特有的接口，专门用于在请求数据绑定过程中处理自定义类型的转换。
*/
