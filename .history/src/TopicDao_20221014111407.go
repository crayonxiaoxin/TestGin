package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopicDetail(ctx *gin.Context) {
	ctx.String(http.StatusOK, "获取帖子ID为 %v", ctx.Param("topic_id"))
}

func GetTopicList(ctx *gin.Context) {
	if ctx.Query("username") == "" {
		ctx.String(http.StatusOK, "获取列表")
	} else {
		ctx.String(http.StatusOK, "获取名为 %v 的帖子", ctx.Query("username"))
	}
}

func NewTopic(ctx *gin.Context) {
	ctx.String(http.StatusOK, "新增帖子")
}

func DeleteTopic(ctx *gin.Context) {
	ctx.String(http.StatusOK, "删除帖子 %v", ctx.Param("topic_id"))
}
