package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error500MW(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			ctx.HTML(http.StatusOK, "common/500", gin.H{ // use definition name in 500.html
				"code":  500,
				"error": err,
			})
		}
	}()

	ctx.Next()
}
