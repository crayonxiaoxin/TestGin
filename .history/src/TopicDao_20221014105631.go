package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopicDetail(ctx *gin.Context) {
	ctx.String(http.StatusOK, "获取帖子ID为 %s", ctx.Param("topic_id"))
}
