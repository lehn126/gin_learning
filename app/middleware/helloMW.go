package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func ResponseTimeMW(ctx *gin.Context) {
	timeBefore := time.Now()
	log.Println("request from", ctx.Request.RequestURI, "start at", timeBefore.Format(time.DateTime))
	ctx.Next()

	timeAfter := time.Now()
	dur := timeAfter.Sub(timeBefore).Seconds()
	log.Println("request from", ctx.Request.RequestURI, "end at", timeAfter.Format(time.DateTime), ", during:", dur, "s")
}
