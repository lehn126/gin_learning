package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error404MW(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "common/404", nil) // use definition name in 404.html
}
