package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/topic/:topic_id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "获取帖子ID为 %s", ctx.Param("topic_id"))
	})

	router.Run(":8082")
}
