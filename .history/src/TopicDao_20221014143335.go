package src

import (
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
	query := TopicQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, query)
	}
}

func NewTopic(ctx *gin.Context) { // 新增一条帖子
	// 需要登录
	// ctx.String(http.StatusOK, "新增帖子")
	topic := Topic{}
	err := ctx.BindQuery(&topic)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, topic)
	}
}

func NewTopics(ctx *gin.Context) { // 新增多条帖子
	// 需要登录
	topic := Topics{}
	err := ctx.BindQuery(&topic)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, topic)
	}
}

func DeleteTopic(ctx *gin.Context) {
	// 需要登录
	ctx.String(http.StatusOK, "删除帖子 %s", ctx.Param("topic_id"))
}
