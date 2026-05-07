package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	避免记录查询字符串
		查询字符串通常包含敏感信息，如 API 令牌、密码、会话 ID 或个人身份信息（PII）。
		记录这些值会产生安全风险，并可能违反 GDPR 或 HIPAA 等隐私法规。
		通过从日志中剥离查询字符串，可以减少通过日志文件、监控系统或错误报告工具泄露敏感数据的风险

		使用 LoggerConfig 中的 SkipQueryString 选项可以防止查询字符串出现在日志中。
		启用后，对 /path?token=secret&user=alice 的请求将简单地记录为 /path

		使用 SkipQueryString 要禁用 自定义日志 和 gin.Logger()中间件 否则日志中还是展示敏感信息
*/

func AvoidLoggingQueryStrings(c *gin.Context) {
	q := c.Query("q")
	t := c.Query("token")
	c.String(http.StatusOK, "searching for: "+q+t)
	/* c.JSON(http.StatusOK, gin.H{
		"success": "searching for: " + q,
	}) */
}
