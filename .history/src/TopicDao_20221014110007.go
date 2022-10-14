package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopicDetail(ctx *gin.Context) {
	ctx.String(http.StatusOK, "获取帖子ID为 %s", ctx.Param("topic_id"))
}

func GetTopicList(ctx *gin.Context) {
	if ctx.Query("username") == "" {
		ctx.String(http.StatusOK, "获取列表")
	} else {
		ctx.String(http.StatusOK, "获取名为 %s 的帖子", ctx.Query("username"))
	}
}
