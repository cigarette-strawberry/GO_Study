package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
	中间件中的 Goroutine
		在中间件或处理函数中启动新的 Goroutine 时，不应该在其中使用原始上下文，必须使用只读副本

	为什么 c.Copy() 至关重要
		Gin 使用 sync.Pool 来复用 gin.Context 对象以提高性能。一旦处理函数返回，gin.Context 就会被返回到池中，可能会被分配给一个完全不同的请求。
		如果此时一个 goroutine 仍然持有对原始上下文的引用，它将读取或写入现在属于另一个请求的字段。这会导致竞态条件、数据损坏或 panic

		调用 c.Copy() 会创建一个上下文快照，可以在处理函数返回后安全使用。副本包含请求、URL、键和其他只读数据，但与池的生命周期分离
*/

func GoroutinesInsideAMiddleware(c *gin.Context) {
	// 创建一个副本，以便在协程内部使用
	cCp := c.Copy()
	go func() {
		// 通过调用 `time.Sleep()` 来模拟一个耗时的任务。5 秒钟
		time.Sleep(5 * time.Second)

		// 请注意，您正在使用复制的上下文“cCp”，非常重要
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()

	c.JSON(http.StatusOK, gin.H{
		"success": "123",
	})

	// 通过调用 `time.Sleep()` 来模拟一个耗时的任务。5 秒钟
	time.Sleep(5 * time.Second)

	// 由于我们并未使用协程，所以无需复制上下文
	log.Println("Done! in path " + c.Request.URL.Path)

}
