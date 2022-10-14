package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"

	. "crayonxiaoxin/test_gin/src" // 前面加 . 后， src.xxx 可以直接写成 xxx（即省略src）
)

// var db *gorm.DB

// func init() {
// 	dsn := "root:New4you!@tcp(127.0.0.1:3306)/topics?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("连接数据库失败")
// 	} else {
// 		log.Println("连接数据库成功")
// 		db = d
// 	}
// }

func main() {
	router := gin.Default()

	// 1. 一般
	// // http://localhost:8082/topic/123
	// router.GET("/topic/:topic_id", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "获取帖子ID为 %s", ctx.Param("topic_id"))
	// })

	// // http://localhost:8082/topic/top
	// router.GET("/topic/top", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "获取帖子 top")
	// })

	// 2. 版本控制
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

	// // 3. 路由分组
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

	// 4.
	// 自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topicurl", TopicUrl)
		v.RegisterValidation("topicsize", TopicSize)
	}

	// 路由分组
	v1 := router.Group("/v1/topics") // 单条帖子
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

	v2 := router.Group("/v1/mtopics") // 多条帖子
	{
		v2.Use(NeedLogin())
		{
			v2.POST("", NewTopics)
		}
	}

	router.Run(":8082")
}
