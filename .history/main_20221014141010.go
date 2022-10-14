package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	. "crayonxiaoxin/test_gin/src" // 前面加 . 后， src.xxx 可以直接写成 xxx（即省略src）
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

	// // 路由分组
	// v1 := router.Group("/v1/topics")
	// // 代码块
	// {
	// 	v1.GET("/:topic_id", func(ctx *gin.Context) {
	// 		ctx.String(http.StatusOK, "获取帖子ID为 %s", ctx.Param("topic_id"))
	// 	})
	// 	v1.GET("", func(ctx *gin.Context) {
	// 		if ctx.Query("username") == "" {
	// 			ctx.String(http.StatusOK, "获取列表")
	// 		} else {
	// 			ctx.String(http.StatusOK, "获取名为 %s 的帖子", ctx.Query("username"))
	// 		}
	// 	})
	// }

	// 自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", TopicUrl)
	}

	// 路由分组
	v1 := router.Group("/v1/topics")
	// 代码块
	{
		v1.GET("/:topic_id", GetTopicDetail)
		v1.GET("", GetTopicList)

		v1.Use(NeedLogin()) //需要登录
		{
			v1.POST("", NewTopic)
			v1.DELETE("/:topic_id", DeleteTopic)
		}
	}

	router.Run(":8082")
}
