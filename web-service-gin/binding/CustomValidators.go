package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

/*
	除了内置的验证器（如 required、email、min、max），你还可以注册自己的自定义验证函数
  如果你的时间字符串总是以 2006 - 01 - 02 这样的标准格式出现，将 time_format 改为 "2006 - 01 - 02"，这样就可以正确解析常见的 ISO 8601 格式的时间字符串。
*/

// gtfield = CheckIn 是一个自定义的验证规则，表示 CheckOut 的值必须大于 CheckIn 的值
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

// 声明了一个名为 bookableDate 的变量，其类型为 validator.Func
// 这是 bookableDate 变量所指向的具体函数。该函数接收一个 validator.FieldLevel 类型的参数 fl，返回一个布尔值
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		// 检查当前时间是否在 date（即被验证的时间字段值）之后。如果是，说明该日期已经过去，不符合可预订条件，返回 false
		if today.After(date) {
			return false
		}
	}
	return true
}

func CustomValidators(c *gin.Context) {
	// binding.Validator.Engine()：获取 Gin 使用的验证器的底层引擎
	// 尝试将验证器引擎断言为 *validator.Validate 类型（假设使用的是 go - playground/validator/v10 库）。如果断言成功，v 会得到指向验证器实例的指针，ok 为 true；否则，ok 为 false
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 在验证器实例 v 上注册名为 bookabledate 的自定义验证规则，该规则对应的验证函数是前面定义的 bookableDate 函数。这样，当在结构体标签中使用 bookabledate 验证规则时，就会调用这个函数进行验证
		v.RegisterValidation("bookabledate", bookableDate)
	}

	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
}
