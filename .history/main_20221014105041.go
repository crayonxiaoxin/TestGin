package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 一般
	// // http://localhost:8082/topic/123
	// router.GET("/topic/:topic_id", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "获取帖子ID为 %s", ctx.Param("topic_id"))
	// })

	// // http://localhost:8082/topic/top
	// router.GET("/topic/top", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "获取帖子 top")
	// })

	// 版本控制
	// // http://localhost:8082/v1/topics/123
	// router.GET("/v1/topics/:topic_id", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "获取帖子ID为 %s", ctx.Param("topic_id"))
	// })

	// // http://localhost:8082/v1/topics?username=Tom
	// router.GET("/v1/topics", func(ctx *gin.Context) {
	// 	if ctx.Query("username") == "" {
	// 		ctx.String(http.StatusOK, "获取列表")
	// 	} else {
	// 		ctx.String(http.StatusOK, "获取名为 %s 的帖子", ctx.Query("username"))
	// 	}
	// })

	// 路由分组
	v1 := router.Group("/v1/topics")
	// 代码块
	{
		v1.GET("/:topic_id", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "获取帖子ID为 %s", ctx.Param("topic_id"))
		})
		v1.GET("", func(ctx *gin.Context) {
			if ctx.Query("username") == "" {
				ctx.String(http.StatusOK, "获取列表")
			} else {
				ctx.String(http.StatusOK, "获取名为 %s 的帖子", ctx.Query("username"))
			}
		})
	}

	router.Run(":8082")
}