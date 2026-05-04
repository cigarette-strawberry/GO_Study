package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Gin 提供了多种方法来向客户端提供文件。每种方法适用于不同的用例

c.File(path) — 从本地文件系统提供文件。内容类型会自动检测。当你在编译时已知确切的文件路径或已验证过路径时使用
c.FileFromFS(path, fs) — 从 http.FileSystem 接口提供文件。适用于从嵌入式文件系统（embed.FS）、自定义存储后端提供文件，或当你想限制对特定目录树的访问时
c.FileAttachment(path, filename) — 通过设置 Content-Disposition: attachment 头将文件作为下载提供。
浏览器会提示用户使用你提供的文件名保存文件，而不管磁盘上的原始文件名

如需从 io.Reader 流式传输数据（如远程 URL 或动态生成的内容），请改用 c.DataFromReader()

永远不要将用户输入直接传递给 c.File() 或 c.FileAttachment()。
攻击者可以提供 ../../etc/passwd 之类的路径来读取服务器上的任意文件。
始终验证和清理文件路径，或使用带有受限 http.FileSystem 的 c.FileFromFS() 来将访问限制在特定目录中
*/

// 直接内联显示一个文件（在浏览器中显示）
func LocalFile(c *gin.Context) {
	c.File("static/text.txt")
}

var fs http.FileSystem = http.Dir("static")

// 从 HTTP 文件系统中提供一个文件
func FsFile(c *gin.Context) {
	c.FileFromFS("text.txt", fs)
}

// 将文件作为可下载附件提供，并设置自定义文件名
func Download(c *gin.Context) {
	c.FileAttachment("static/text.txt", "cigarette.txt")
}
