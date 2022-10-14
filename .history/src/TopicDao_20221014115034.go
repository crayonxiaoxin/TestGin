package src

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NeedLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, status := ctx.GetQuery("token"); !status {
			ctx.String(http.StatusUnauthorized, "缺少token参数")
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}

func GetTopicDetail(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "获取帖子ID为 %s", ctx.Param("topic_id"))
	ctx.JSON(http.StatusOK, CreateTopic(101, "帖子标题"))
}

func GetTopicList(ctx *gin.Context) {
	// if ctx.Query("username") == "" {
	// 	ctx.String(http.StatusOK, "获取列表")
	// } else {
	// 	ctx.String(http.StatusOK, "获取名为 %s 的帖子", ctx.Query("username"))
	// }
	tq := TopicQuery{}
	err := ctx.BindQuery(&tq)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func NewTopic(ctx *gin.Context) {
	// 需要登录
	ctx.String(http.StatusOK, "新增帖子")
}

func DeleteTopic(ctx *gin.Context) {
	// 需要登录
	ctx.String(http.StatusOK, "删除帖子 %s", ctx.Param("topic_id"))
}
