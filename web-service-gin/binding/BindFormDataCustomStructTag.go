package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
	默认情况下，Gin 使用 form 结构体标签来绑定表单数据。当你需要绑定一个使用不同标签的结构体时——例如一个你无法修改的外部类型——你可以创建一个自定义绑定来读取你自己的标签。

	当集成第三方库时，如果其结构体使用 url、query 或其他自定义名称而非 form 标签，这将非常有用。

	自定义绑定实现了 binding.Binding 接口，该接口需要一个 Name() string 方法和一个 Bind(*http.Request, any) error 方法。
	binding.MapFormWithTag 辅助函数完成了使用自定义标签将表单值映射到结构体字段的实际工作。
*/

/*
这个常量在后续代码中用于指定结构体标签名，以便从请求表单数据中匹配对应的字段。
32 << 20 表示将 32 左移 20 位，即 32 * 2^20，结果是 33554432 字节，也就是 32MB
*/
const (
	customerTag   = "url"
	defaultMemory = 32 << 20
)

// 定义了一个名为 customerBinding 的结构体，它将用于实现自定义的绑定逻辑。这里结构体为空，因为所有的绑定逻辑将通过方法来实现。
type customerBinding struct{}

// 这是 customerBinding 结构体的方法，用于返回该绑定器的名称。这里返回 "form"，表示这个绑定器主要用于处理表单数据
func (customerBinding) Name() string {
	return "form"
}

func (customerBinding) Bind(req *http.Request, obj any) error {
	// 首先尝试解析请求中的表单数据。ParseForm 方法会解析 application/x - www - form - urlencoded 格式的表单数据。如果解析失败，直接返回错误。
	if err := req.ParseForm(); err != nil {
		return err
	}
	// 接着尝试解析 multipart/form - data 格式的表单数据，使用之前定义的 defaultMemory 作为最大内存限制。如果解析过程中出现错误，且错误不是 http.ErrNotMultipart（表示请求不是多部分表单时的错误），则返回错误。如果是 http.ErrNotMultipart 错误，说明请求不是多部分表单，这在这种情况下是可以接受的，继续执行后续代码。
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	// 调用 binding.MapFormWithTag 函数，将解析后的表单数据（req.Form）与目标结构体（obj）进行映射，使用 customerTag（即 "url"）作为结构体标签来匹配字段
	if err := binding.MapFormWithTag(obj, req.Form, customerTag); err != nil {
		return err
	}
	return validate(obj)
}

// 检查是否有验证器（binding.Validator）被设置。如果没有设置验证器，直接返回 nil，表示验证通过
func validate(obj any) error {
	if binding.Validator == nil {
		return nil
	}
	// 如果设置了验证器，调用验证器的 ValidateStruct 方法对目标结构体（obj）进行验证，并返回验证结果。
	return binding.Validator.ValidateStruct(obj)
}

// 定义了一个名为 FormA 的结构体，它有一个字段 FieldA，使用 url:"field_a" 标签，这个标签将用于与自定义绑定逻辑中的 customerTag（即 "url"）匹配，以便从表单数据中提取对应的值。
// FormA is an external type that we can't modify its tag
type FormA struct {
	FieldA string `url:"field_a"`
}

func BindFormDataCustomStructTag(c *gin.Context) {
	var urlBinding = customerBinding{}
	var opt FormA
	if err := c.MustBindWith(&opt, urlBinding); err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"field_a": opt.FieldA})
}
